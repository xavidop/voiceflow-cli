package voiceflow

import (
	"os"
	"strings"

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

// SetVoiceflowURLOverrides populates custom URL global variables from
// environment variables when they have not been set via CLI flags.
func SetVoiceflowURLOverrides() {
	if global.VoiceflowSubdomain == "" {
		global.VoiceflowSubdomain = os.Getenv("VF_SUBDOMAIN")
	}
	if global.VoiceflowAPIURL == "" {
		global.VoiceflowAPIURL = strings.TrimRight(os.Getenv("VF_API_URL"), "/")
	}
	if global.VoiceflowRuntimeURL == "" {
		global.VoiceflowRuntimeURL = strings.TrimRight(os.Getenv("VF_RUNTIME_URL"), "/")
	}
	if global.VoiceflowAnalyticsURL == "" {
		global.VoiceflowAnalyticsURL = strings.TrimRight(os.Getenv("VF_ANALYTICS_URL"), "/")
	}
}
