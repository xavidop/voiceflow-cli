package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/openai"
	"github.com/xavidop/voiceflow-cli/internal/voiceflow"
	"github.com/xavidop/voiceflow-cli/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Voiceflow CLI API server",
	Long: `Start the Voiceflow CLI API server to expose test execution endpoints.

The server provides HTTP endpoints for:
- Executing test suites
- Checking test execution status
- Retrieving system information

The server includes auto-generated OpenAPI/Swagger documentation available at /swagger/index.html`,
	Example: `  # Start server on default port (8080)
  voiceflow server

  # Start server on custom port
  voiceflow server --port 9090

  # Start server with debug mode
  voiceflow server --debug

  # Start server with custom host
  voiceflow server --host 127.0.0.1 --port 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		host, _ := cmd.Flags().GetString("host")
		debug, _ := cmd.Flags().GetBool("debug")
		corsEnabled, _ := cmd.Flags().GetBool("cors")
		swaggerEnabled, _ := cmd.Flags().GetBool("swagger")

		voiceflow.SetVoiceflowAPIKey()
		openai.SetOpenAIAPIKey()

		config := &server.ServerConfig{
			Port:           port,
			Host:           host,
			Debug:          debug,
			CORSEnabled:    corsEnabled,
			SwaggerEnabled: swaggerEnabled,
		}

		srv := server.NewServer(config)
		if err := srv.Start(); err != nil {
			global.Log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("port", "p", "8080", "Port to run the server on")
	serverCmd.Flags().StringP("host", "H", "0.0.0.0", "Host to bind the server to")
	serverCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")
	serverCmd.Flags().Bool("cors", true, "Enable CORS middleware")
	serverCmd.Flags().Bool("swagger", true, "Enable Swagger documentation endpoint")
}
