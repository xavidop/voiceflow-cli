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
	return FetchStateWithOverrides(EnvironmentName, userID, "", "")
}

func FetchStateWithOverrides(EnvironmentName, userID, apiKeyOverride, subdomainOverride string) (state.State, error) {
	// Use the provided subdomain override, or fall back to global if not provided
	subdomain := global.VoiceflowSubdomain
	if subdomainOverride != "" {
		subdomain = subdomainOverride
	}

	// Add the dot prefix if subdomain is not empty
	if subdomain != "" {
		subdomain = "." + subdomain
	}

	url := fmt.Sprintf("https://general-runtime%s.voiceflow.com/state/user/%s", subdomain, userID)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return state.State{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	// Use the provided API key override, or fall back to global if not provided
	apiKey := global.VoiceflowAPIKey
	if apiKeyOverride != "" {
		apiKey = apiKeyOverride
	}
	req.Header.Add("Authorization", apiKey)
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
