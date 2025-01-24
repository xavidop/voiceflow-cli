package openai

import (
	"os"

	"github.com/xavidop/voiceflow-cli/internal/global"
)

func SetOpenAIAPIKey() {

	if global.OpenAIAPIKey == "" {
		global.OpenAIAPIKey = os.Getenv("OPENAI_API_KEY")
	}
}
