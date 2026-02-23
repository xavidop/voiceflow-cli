package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/pkg/test"
)

// --- WebSocket message protocol ---

// WSIncomingMessage is the envelope for all client-to-server messages.
type WSIncomingMessage struct {
	Action string          `json:"action"` // "execute", "cancel", "status"
	ID     string          `json:"id,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

// WSOutgoingMessage is the envelope for all server-to-client messages.
type WSOutgoingMessage struct {
	Type    string      `json:"type"` // "log", "status", "result", "error"
	ID      string      `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// --- Upgrader ---

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (CORS is already handled by Gin middleware)
	},
}

// --- WebSocket connection wrapper ---

// wsConn wraps a websocket.Conn with a write mutex so concurrent goroutines
// can safely send messages.
type wsConn struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func (wc *wsConn) writeJSON(v interface{}) error {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.conn.WriteJSON(v)
}

// --- Handler ---

// HandleWebSocket upgrades an HTTP connection to WebSocket and processes
// messages using the same business logic as the REST endpoints.
//
// Protocol:
//
//	Client → Server (JSON):
//	  { "action": "execute", "data": <TestExecutionRequest> }
//	  { "action": "cancel",  "id": "<execution-id>" }
//	  { "action": "status",  "id": "<execution-id>" }
//
//	Server → Client (JSON):
//	  { "type": "log",    "id": "...", "message": "..." }       – real-time log line
//	  { "type": "status", "id": "...", "data": <TestStatusResponse> }
//	  { "type": "result", "id": "...", "data": <TestStatusResponse> }  – final result
//	  { "type": "error",  "message": "..." }                    – protocol / validation error
func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Log.Errorf("WebSocket upgrade failed: %v", err)
		return
	}
	defer func() { _ = ws.Close() }()

	conn := &wsConn{conn: ws}

	global.Log.Info("WebSocket client connected")

	for {
		_, raw, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				global.Log.Errorf("WebSocket read error: %v", err)
			}
			break
		}

		var msg WSIncomingMessage
		if err := json.Unmarshal(raw, &msg); err != nil {
			_ = conn.writeJSON(WSOutgoingMessage{
				Type:    "error",
				Message: "Invalid JSON: " + err.Error(),
			})
			continue
		}

		switch msg.Action {
		case "execute":
			handleWSExecute(conn, msg)
		case "cancel":
			handleWSCancel(conn, msg)
		case "status":
			handleWSStatus(conn, msg)
		default:
			_ = conn.writeJSON(WSOutgoingMessage{
				Type:    "error",
				Message: "Unknown action: " + msg.Action + ". Supported actions: execute, cancel, status",
			})
		}
	}

	global.Log.Info("WebSocket client disconnected")
}

// handleWSExecute starts a test suite execution and streams logs back over
// the WebSocket connection in real-time, reusing the same business logic as
// the REST handler.
func handleWSExecute(conn *wsConn, msg WSIncomingMessage) {
	var req TestExecutionRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		_ = conn.writeJSON(WSOutgoingMessage{
			Type:    "error",
			Message: "Invalid execute payload: " + err.Error(),
		})
		return
	}

	executionID := uuid.New().String()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Create execution record (shared with REST status/cancel endpoints)
	execution := &TestExecution{
		ID:        executionID,
		Status:    "running",
		StartedAt: time.Now(),
		Logs:      []string{},
	}
	execution.cancelFunc = cancel

	testExecutionsMutex.Lock()
	testExecutions[executionID] = execution
	testExecutionsMutex.Unlock()

	// Notify client that execution has started
	_ = conn.writeJSON(WSOutgoingMessage{
		Type:    "status",
		ID:      executionID,
		Message: "running",
		Data: TestExecutionResponse{
			ID:        executionID,
			Status:    "running",
			StartedAt: execution.StartedAt,
			Logs:      []string{"Test execution started"},
		},
	})

	// Build the suite payload (same conversion as REST handler)
	httpSuite := test.HTTPSuiteRequest{
		Name:               req.Suite.Name,
		Description:        req.Suite.Description,
		EnvironmentName:    req.Suite.EnvironmentName,
		NewSessionPerTest:  req.Suite.NewSessionPerTest,
		Tests:              make([]test.HTTPTestRequest, len(req.Suite.Tests)),
		ApiKey:             req.ApiKey,
		VoiceflowSubdomain: req.VoiceflowSubdomain,
	}
	for i, t := range req.Suite.Tests {
		httpSuite.Tests[i] = test.HTTPTestRequest{
			ID:   t.ID,
			Test: t.Test,
		}
	}

	// Run in a goroutine so the WebSocket read loop continues to process
	// cancel/status messages while tests execute.
	go func() {
		defer cancel()

		execution.AddLog("Starting test suite execution...")
		execution.AddLog("Suite: " + req.Suite.Name)
		execution.AddLog("Environment: " + req.Suite.EnvironmentName)

		// Stream each log line to the WebSocket client in real-time
		onLog := func(message string) {
			_ = conn.writeJSON(WSOutgoingMessage{
				Type:    "log",
				ID:      executionID,
				Message: message,
			})
		}

		result := test.ExecuteFromHTTPRequestWithCallback(ctx, httpSuite, onLog)

		// Persist logs in the execution record
		for _, logLine := range result.Logs {
			execution.AddLog(logLine)
		}

		if !result.Success {
			if ctx.Err() != nil {
				execution.AddLog("Test suite execution cancelled")
			} else {
				execution.SetError(result.Error.Error())
				execution.AddLog("Test suite execution failed: " + result.Error.Error())
			}
		} else {
			execution.SetStatus("completed")
			execution.AddLog("Test suite execution completed successfully")
		}

		// Send final result
		_ = conn.writeJSON(WSOutgoingMessage{
			Type: "result",
			ID:   executionID,
			Data: execution.GetStatus(),
		})
	}()
}

// handleWSCancel cancels a running test execution (reuses the same shared
// TestExecution store as the REST handler).
func handleWSCancel(conn *wsConn, msg WSIncomingMessage) {
	if msg.ID == "" {
		_ = conn.writeJSON(WSOutgoingMessage{
			Type:    "error",
			Message: "Missing execution id for cancel action",
		})
		return
	}

	testExecutionsMutex.RLock()
	execution, exists := testExecutions[msg.ID]
	testExecutionsMutex.RUnlock()

	if !exists {
		_ = conn.writeJSON(WSOutgoingMessage{
			Type:    "error",
			ID:      msg.ID,
			Message: "Test execution not found: " + msg.ID,
		})
		return
	}

	if !execution.Cancel() {
		_ = conn.writeJSON(WSOutgoingMessage{
			Type:    "error",
			ID:      msg.ID,
			Message: "Cannot cancel: execution is not running (current status: " + execution.GetStatus().Status + ")",
		})
		return
	}

	_ = conn.writeJSON(WSOutgoingMessage{
		Type: "status",
		ID:   msg.ID,
		Data: execution.GetStatus(),
	})
}

// handleWSStatus returns the current status of an execution.
func handleWSStatus(conn *wsConn, msg WSIncomingMessage) {
	if msg.ID == "" {
		_ = conn.writeJSON(WSOutgoingMessage{
			Type:    "error",
			Message: "Missing execution id for status action",
		})
		return
	}

	testExecutionsMutex.RLock()
	execution, exists := testExecutions[msg.ID]
	testExecutionsMutex.RUnlock()

	if !exists {
		_ = conn.writeJSON(WSOutgoingMessage{
			Type:    "error",
			ID:      msg.ID,
			Message: "Test execution not found: " + msg.ID,
		})
		return
	}

	_ = conn.writeJSON(WSOutgoingMessage{
		Type: "status",
		ID:   msg.ID,
		Data: execution.GetStatus(),
	})
}
