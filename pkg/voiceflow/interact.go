package voiceflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/interact"
)

func CallInteractionAPI(EnvironmentName, userID string, interaction tests.Interaction) ([]interact.InteractionResponse, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://general-runtime%s.voiceflow.com/state/user/%s/interact?logs=off", global.VoiceflowSubdomain, userID)
	var interatctionRequest interact.InteratctionRequest
	switch interaction.User.Type {
	case "launch":
		interatctionRequest = interact.InteratctionRequest{
			Action: interact.Action{
				Type: "launch",
			},
		}
	case "text":
		interatctionRequest = interact.InteratctionRequest{
			Action: interact.Action{
				Type:    "text",
				Payload: interaction.User.Text,
			},
		}
	}
	byts, err := json.Marshal(interatctionRequest)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error marshalling request: %v", err)
	}

	payload := strings.NewReader(string(byts))

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)
	req.Header.Add("versionID", EnvironmentName)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error calling API: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error reading response: %v", err)
	}

	var interactions []interact.InteractionResponse
	err = json.Unmarshal([]byte(string(body)), &interactions)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return interactions, nil
}
