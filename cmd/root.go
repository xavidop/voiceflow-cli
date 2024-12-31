package cmd

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/cmd/cmdutils"
	test "github.com/xavidop/voiceflow-cli/cmd/test"
	"github.com/xavidop/voiceflow-cli/internal/global"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "voiceflow",
	Short: "Voiceflow CLI",
	Long: `Welcome to voiceflow-cli!

This utility provides you with an easy way to interact
with your Voiceflow agents.

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

func init() {
	if err := godotenv.Load(); err != nil {
		global.Log.Debug("No .env file found")
	}
	// Add the subcommands
	test.Register(rootCmd)

	// Add the subcommands
	rootCmd.PersistentFlags().BoolVarP(&global.Verbose, "verbose", "v", false, "verbose error output (with stack trace) (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.VoiceflowAPIKey, "voiceflow-api-key", "x", "", "Voiceflow API Key (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.VoiceflowBaseURL, "voiceflow-base-url", "b", "https://general-runtime.voiceflow.com", "Voiceflow Base URL (optional). Default: https://general-runtime.voiceflow.com")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipUpdate, "skip-update-check", "u", false, "Skip the check for updates check run before every command (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.Output, "output-format", "o", "text", "Output Format. Options: text, json. Default: text (optional)")

}
