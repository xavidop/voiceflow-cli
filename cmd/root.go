package cmd

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/agent"
	"github.com/xavidop/voiceflow-cli/cmd/analytics"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	"github.com/xavidop/voiceflow-cli/cmd/document"
	"github.com/xavidop/voiceflow-cli/cmd/kb"
	test "github.com/xavidop/voiceflow-cli/cmd/test"
	"github.com/xavidop/voiceflow-cli/cmd/transcript"
	"github.com/xavidop/voiceflow-cli/internal/global"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "voiceflow",
	Short: "Voiceflow CLI",
	Long: `Welcome to voiceflow-cli!

This utility provides you with an easy way to interact
with your Voiceflow agents.

You can run it in two modes:
1. CLI mode (default): Execute commands directly
2. Server mode: Run as a local HTTP server with REST API endpoints

For server mode, use --server flag or set SERVER=true in your .env file.
Example: voiceflow --server --port 3000

You can find the documentation at https://github.com/xavidop/voiceflow-cli.

Please file all bug reports on GitHub at https://github.com/xavidop/voiceflow-cli/issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				os.Exit(1)
			}
			os.Exit(0)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		global.Log.Error(errors.Errorf("%s", err))
		os.Exit(1)
	}
}

// RegisterServerFlags registers the server-related flags
func RegisterServerFlags() {
	// Add server flags
	rootCmd.Flags().BoolP("server", "s", false, "Run in server mode with REST API endpoints")
	rootCmd.Flags().StringP("port", "p", "8080", "Port to run the server on (only used with --server)")
}

func init() {
	if err := godotenv.Load(); err != nil {
		global.Log.Debug("No .env file found")
	}
	// Add the subcommands
	test.Register(rootCmd)
	transcript.Register(rootCmd)
	analytics.Register(rootCmd)
	agent.Register(rootCmd)
	kb.Register(rootCmd)
	document.Register(rootCmd)

	// Add the flags
	rootCmd.PersistentFlags().BoolVarP(&global.Verbose, "verbose", "v", false, "verbose error output (with stack trace) (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.VoiceflowAPIKey, "voiceflow-api-key", "x", "", "Voiceflow API Key (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.VoiceflowSubdomain, "voiceflow-subdomain", "b", "", "Voiceflow Base URL (optional). Default: empty")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipUpdate, "skip-update-check", "u", false, "Skip the check for updates check run before every command (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.Output, "output-format", "o", "text", "Output Format. Options: text, json. Default: text (optional)")
}
