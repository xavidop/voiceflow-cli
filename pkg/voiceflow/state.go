package voiceflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/state"
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

func FetchState(EnvironmentName, userID string) (state.State, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://general-runtime%s.voiceflow.com/state/user/%s", global.VoiceflowSubdomain, userID)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return state.State{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)
	req.Header.Add("versionID", EnvironmentName)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return state.State{}, fmt.Errorf("error calling API: %v", err)
	}
	defer utils.SafeClose(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return state.State{}, fmt.Errorf("error reading response: %v", err)
	}

	var stateResponse state.State
	err = json.Unmarshal([]byte(string(body)), &stateResponse)
	if err != nil {
		return state.State{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return stateResponse, nil
}
