package agent

import (
	"fmt"

	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
)

func Export(agentID, versionID, outputFile string) error {

	agent, err := voiceflow.ExportAgent(agentID, versionID)
	if err != nil {
		return fmt.Errorf("failed to fetch agent %s: %w", versionID, err)
	}

	err = voiceflow.SaveAgent(agent, outputFile)
	if err != nil {
		return fmt.Errorf("failed to save agentID %s with versionID %s: %w", agentID, versionID, err)
	}

	return nil
}
