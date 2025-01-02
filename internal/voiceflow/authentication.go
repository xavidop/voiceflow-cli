package voiceflow

import (
	"os"

	"github.com/xavidop/voiceflow-cli/internal/global"
)

func SetVoiceflowAPIKey() {
	if global.VoiceflowAPIKey == "" && os.Getenv("VF_API_KEY") == "" {
		global.Log.Errorf("Voiceflow API Key is required")
		os.Exit(1)
	}
	if global.VoiceflowAPIKey == "" {
		global.VoiceflowAPIKey = os.Getenv("VF_API_KEY")
	}
}
