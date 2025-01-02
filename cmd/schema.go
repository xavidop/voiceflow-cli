package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
)

type schemaCmd struct {
	cmd    *cobra.Command
	output string
}

func newSchemaCmd() *schemaCmd {
	root := &schemaCmd{}
	cmd := &cobra.Command{
		Use:           "jsonschema",
		Aliases:       []string{"schema"},
		Short:         "outputs voiceflow's JSON schema",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := generateTestSchema(root.output); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&root.output, "output-folder", "f", "-", "Where to save the JSONSchema file")
	_ = cmd.Flags().SetAnnotation("output-file", cobra.BashCompFilenameExt, []string{"json"})

	root.cmd = cmd
	return root
}

func generateTestSchema(output string) error {
	suiteSchema := jsonschema.Reflect(&tests.Suite{})
	suiteSchema.Definitions["Tests"] = jsonschema.Reflect(&[]tests.Test{})
	suiteSchema.Description = "voiceflow-cli suite definition file"

	testSchema := jsonschema.Reflect(&tests.Test{})
	testSchema.Description = "voiceflow-cli Conversation Profiler test definition file"

	testBts, err := json.MarshalIndent(testSchema, "	", "	")
	if err != nil {
		return fmt.Errorf("failed to create test jsonschema: %w", err)
	}

	suiteBts, err := json.MarshalIndent(suiteSchema, "	", "	")
	if err != nil {
		return fmt.Errorf("failed to create suite jsonschema: %w", err)
	}
	if output == "-" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}

	if err := os.WriteFile(output+"/conversationsuite.json", suiteBts, 0o666); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}
	if err := os.WriteFile(output+"/conversationtest.json", testBts, 0o666); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(newSchemaCmd().cmd)
}
