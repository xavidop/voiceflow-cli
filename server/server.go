package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xavidop/voiceflow-cli/internal/global"
	_ "github.com/xavidop/voiceflow-cli/server/docs" // Import swagger docs
	"github.com/xavidop/voiceflow-cli/server/handlers"
)

// @title Voiceflow CLI API
// @version 1.0
// @description API server for Voiceflow CLI test execution and management
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// ServerConfig holds the server configuration
type ServerConfig struct {
	Port           string
	Host           string
	Debug          bool
	CORSEnabled    bool
	SwaggerEnabled bool
}

// Server represents the HTTP server
type Server struct {
	config *ServerConfig
	router *gin.Engine
	server *http.Server
}

// NewServer creates a new server instance
func NewServer(config *ServerConfig) *Server {
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if config.CORSEnabled {
		router.Use(cors.Default())
	}

	return &Server{
		config: config,
		router: router,
	}
}

// setupRoutes configures the API routes
func (s *Server) setupRoutes() {
	// Health check endpoint
	s.router.GET("/health", handlers.HealthCheck)

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Test endpoints
		tests := v1.Group("/tests")
		{
			tests.POST("/execute", handlers.ExecuteTestSuite)
			tests.GET("/status/:id", handlers.GetTestStatus)
		}

		// System endpoints
		system := v1.Group("/system")
		{
			system.GET("/info", handlers.GetSystemInfo)
		}
	}

	// Swagger documentation
	if s.config.SwaggerEnabled {
		s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.setupRoutes()

	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)

	s.server = &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	global.Log.Infof("Starting Voiceflow CLI API server on %s", addr)
	if s.config.SwaggerEnabled {
		global.Log.Infof("Swagger documentation available at http://%s/swagger/index.html", addr)
	}

	// Start server in a goroutine
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return s.waitForShutdown()
}

// waitForShutdown waits for interrupt signal and gracefully shuts down the server
func (s *Server) waitForShutdown() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		global.Log.Errorf("Server forced to shutdown: %v", err)
		return err
	}

	global.Log.Info("Server exited")
	return nil
}

// DefaultConfig returns a default server configuration
func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		Port:           "8080",
		Host:           "0.0.0.0",
		Debug:          false,
		CORSEnabled:    true,
		SwaggerEnabled: true,
	}
}
