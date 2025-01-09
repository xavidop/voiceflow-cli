package server

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/pkg/document"
	"github.com/xavidop/voiceflow-cli/pkg/transcript"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

type Server struct {
	router *gin.Engine
}

// authMiddleware handles API key authorization
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Authorization header first
		apiKey := c.GetHeader("Authorization")
		// If still not found, use the one from .env
		if apiKey == "" {
			apiKey = os.Getenv("VOICEFLOW_API_KEY")
		}

		if apiKey == "" {
			c.JSON(401, gin.H{
				"error": "API key is required. Set it in the Authorization header or VOICEFLOW_API_KEY environment variable",
			})
			c.Abort()
			return
		}

		// Store the API key in the context for use in handlers
		global.VoiceflowAPIKey = apiKey
		c.Next()
	}
}

func NewServer() *Server {
	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Configure trusted proxies based on environment
	if proxyList := os.Getenv("TRUSTED_PROXIES"); proxyList != "" {
		router.SetTrustedProxies(strings.Split(proxyList, ","))
	} else {
		// Default to only trusting localhost
		router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	}

	// Add the auth middleware to all routes
	router.Use(authMiddleware())

	server := &Server{
		router: router,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Agent routes
	agentGroup := s.router.Group("/api/agent")
	{
		agentGroup.GET("/export", s.handleAgentExport)
		// Add other agent endpoints
	}

	// Analytics routes
	analyticsGroup := s.router.Group("/api/analytics")
	{
		analyticsGroup.GET("/fetch", s.handleAnalyticsFetch)
		// Add other analytics endpoints
	}

	// Document routes
	documentGroup := s.router.Group("/api/document")
	{
		documentGroup.GET("/fetch", s.handleDocumentFetch)
		documentGroup.POST("/upload-url", s.handleDocumentUploadURL)
		documentGroup.POST("/upload-file", s.handleDocumentUploadFile)
	}

	// KB routes
	kbGroup := s.router.Group("/api/kb")
	{
		kbGroup.POST("/query", s.handleKBQuery)
		// Add other kb endpoints
	}

	// Transcript routes
	transcriptGroup := s.router.Group("/api/transcript")
	{
		transcriptGroup.GET("/fetch", s.handleTranscriptFetch)
		// Add other transcript endpoints
	}
}

func (s *Server) Run(addr string) error {
	// Get port from command-line parameter first, then environment variable, then default
	port := addr
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
	}

	// Ensure port starts with ":"
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("🚀 Server is running on http://localhost%s\n", port)
	return s.router.Run(port)
}

// Handler implementations
func (s *Server) handleAgentExport(c *gin.Context) {
	agentID := c.Query("agent-id")
	versionID := c.DefaultQuery("version-id", "development")

	if agentID == "" {
		c.JSON(400, gin.H{"error": "agent-id is required"})
		return
	}

	// Get the agent data directly
	agentData, err := voiceflow.ExportAgent(agentID, versionID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Parse the agent data to return proper JSON
	var result interface{}
	if err := json.Unmarshal([]byte(agentData), &result); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse agent data"})
		return
	}

	c.JSON(200, result)
}

func (s *Server) handleAnalyticsFetch(c *gin.Context) {
	agentID := c.Query("agent-id")
	startTime := c.Query("start-time")
	endTime := c.Query("end-time")
	limitStr := c.DefaultQuery("limit", "100")
	analyticsToFetch := c.QueryArray("analytics")

	if agentID == "" {
		c.JSON(400, gin.H{"error": "agent-id is required"})
		return
	}

	// If no analytics types specified, use all available types
	if len(analyticsToFetch) == 0 {
		analyticsToFetch = []string{
			"interactions",
			"sessions",
			"top_intents",
			"top_slots",
			"understood_messages",
			"unique_users",
			"token_usage",
		}
	} else {
		// Validate analytics types
		validTypes := map[string]bool{
			"interactions":        true,
			"sessions":           true,
			"top_intents":        true,
			"top_slots":          true,
			"understood_messages": true,
			"unique_users":       true,
			"token_usage":        true,
		}

		for _, t := range analyticsToFetch {
			if !validTypes[t] {
				c.JSON(400, gin.H{
					"error": "invalid analytics type: " + t,
					"valid_types": []string{
						"interactions",
						"sessions",
						"top_intents",
						"top_slots",
						"understood_messages",
						"unique_users",
						"token_usage",
					},
				})
				return
			}
		}
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid limit value"})
		return
	}

	// Parse filters
	startTimeDate, endTimeDate, limitInt, err := voiceflow.ParseFilters(startTime, endTime, limit)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Fetch analytics
	analytics, err := voiceflow.FetchAnalytics(agentID, startTimeDate, endTimeDate, limitInt, analyticsToFetch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Parse analytics to return proper JSON
	var result interface{}
	if err := json.Unmarshal([]byte(analytics), &result); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse analytics data"})
		return
	}

	c.JSON(200, result)
}

func (s *Server) handleDocumentFetch(c *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Get other parameters
	documentType := c.Query("document-type")
	includeTags := c.QueryArray("include-tags")
	excludeTags := c.QueryArray("exclude-tags")
	includeAllNonTagged := c.DefaultQuery("include-all-non-tagged", "false") == "true"
	includeAllTagged := c.DefaultQuery("include-all-tagged", "false") == "true"

	// Get the documents data
	documents, err := voiceflow.ListDocuments(
		page,
		limit,
		documentType,
		includeTags,
		excludeTags,
		includeAllNonTagged,
		includeAllTagged,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Parse the documents data to return proper JSON
	var result interface{}
	if err := json.Unmarshal([]byte(documents), &result); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse documents data"})
		return
	}

	c.JSON(200, result)
}

func (s *Server) handleKBQuery(c *gin.Context) {
	var request struct {
		Question           string   `json:"question" binding:"required"`
		Model             string   `json:"model"`
		Temperature       float64  `json:"temperature"`
		ChunkLimit        int      `json:"chunk_limit"`
		Synthesis        bool     `json:"synthesis"`
		SystemPrompt      string   `json:"system_prompt"`
		IncludeTags      []string `json:"include_tags"`
		IncludeOperator  string   `json:"include_operator"`
		ExcludeTags      []string `json:"exclude_tags"`
		ExcludeOperator  string   `json:"exclude_operator"`
		IncludeAllTagged bool     `json:"include_all_tagged"`
		IncludeAllNonTagged bool  `json:"include_all_non_tagged"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Query the knowledge base directly
	response, err := voiceflow.QueryKB(
		request.Question,
		request.Model,
		request.Temperature,
		request.ChunkLimit,
		request.Synthesis,
		request.SystemPrompt,
		request.IncludeTags,
		request.IncludeOperator,
		request.ExcludeTags,
		request.ExcludeOperator,
		request.IncludeAllTagged,
		request.IncludeAllNonTagged,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Parse response to return proper JSON
	var result interface{}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse query results"})
		return
	}

	c.JSON(200, result)
}

func (s *Server) handleTranscriptFetch(c *gin.Context) {
	agentID := c.Query("agent-id")
	transcriptID := c.Query("transcript-id")
	outputDirectory := c.Query("output-directory")

	if err := transcript.Fetch(agentID, transcriptID, outputDirectory); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Transcript fetched successfully"})
}

func (s *Server) handleDocumentUploadURL(c *gin.Context) {
	var request struct {
		URL                    string   `json:"url" binding:"required"`
		Name                   string   `json:"name" binding:"required"`
		Overwrite             bool     `json:"overwrite"`
		MaxChunkSize          int      `json:"max_chunk_size"`
		MarkdownConversion    bool     `json:"markdown_conversion"`
		LLMGeneratedQ         bool     `json:"llm_generated_q"`
		LLMPrependContext     bool     `json:"llm_prepend_context"`
		LLMBasedChunking      bool     `json:"llm_based_chunking"`
		LLMContentSummarization bool   `json:"llm_content_summarization"`
		Tags                  []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := document.UploadURL(
		request.URL,
		request.Name,
		request.Overwrite,
		request.MaxChunkSize,
		request.MarkdownConversion,
		request.LLMGeneratedQ,
		request.LLMPrependContext,
		request.LLMBasedChunking,
		request.LLMContentSummarization,
		request.Tags,
	); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Document uploaded successfully from URL"})
}

func (s *Server) handleDocumentUploadFile(c *gin.Context) {
	// Get query parameters
	overwrite := c.Query("overwrite") == "true"
	maxChunkSize, _ := strconv.Atoi(c.DefaultQuery("maxChunkSize", "1000"))
	markdownConversion := c.Query("markdownConversion") == "true"
	llmGeneratedQ := c.Query("llmGeneratedQ") == "true"
	llmPrependContext := c.Query("llmPrependContext") == "true"
	llmBasedChunking := c.Query("llmBasedChunking") == "true"
	llmContentSummarization := c.Query("llmContentSummarization") == "true"
	tags := c.QueryArray("tags")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	// Create temp file in system's temp directory
	tempFile := filepath.Join(os.TempDir(), "vf_cli_"+file.Filename)
	if err := c.SaveUploadedFile(file, tempFile); err != nil {
		c.JSON(500, gin.H{"error": "failed to save file"})
		return
	}
	defer os.Remove(tempFile) // Clean up temp file after we're done

	// Upload the file to Voiceflow
	response, err := voiceflow.UploadDocumentFile(
		tempFile,
		overwrite,
		maxChunkSize,
		markdownConversion,
		llmGeneratedQ,
		llmPrependContext,
		llmBasedChunking,
		llmContentSummarization,
		tags,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Parse response to return proper JSON
	var result interface{}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse upload response"})
		return
	}

	c.JSON(200, result)
}
