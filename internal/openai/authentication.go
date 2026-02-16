package openai

import (
	"os"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
)

func SetOpenAIAPIKey() {

	if global.OpenAIAPIKey == "" {
		global.OpenAIAPIKey = os.Getenv("OPENAI_API_KEY")
	}
}

func SetOpenAIBaseURL() {
	if global.OpenAIBaseURL == "" {
		global.OpenAIBaseURL = os.Getenv("OPENAI_BASE_URL")
	}

	// Keep current behavior as default while allowing region-specific overrides.
	if global.OpenAIBaseURL == "" {
		global.OpenAIBaseURL = "https://api.openai.com/v1"
	}

	global.OpenAIBaseURL = strings.TrimRight(global.OpenAIBaseURL, "/")
}

func GetChatCompletionsURL() string {
	SetOpenAIBaseURL()
	return global.OpenAIBaseURL + "/chat/completions"
}
