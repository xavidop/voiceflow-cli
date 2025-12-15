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
	"github.com/xavidop/voiceflow-cli/internal/utils"
)

func DialogManagerInteract(environmentName, userID string, interaction tests.Interaction, apiKeyOverride, subdomainOverride string, availableButtons []tests.Button) ([]interact.InteractionResponse, error) {
	// Use the provided subdomain override, or fall back to global if not provided
	subdomain := global.VoiceflowSubdomain
	if subdomainOverride != "" {
		subdomain = subdomainOverride
	}

	// Add the dot prefix if subdomain is not empty
	if subdomain != "" {
		subdomain = "." + subdomain
	}

	url := fmt.Sprintf("https://general-runtime%s.voiceflow.com/state/user/%s/interact?logs=off", subdomain, userID)
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
	case "event":
		interatctionRequest = interact.InteratctionRequest{
			Action: interact.Action{
				Type: "event",
				Payload: interact.EventPayload{
					Event: interact.EventData{
						Name: interaction.User.Event,
					},
				},
			},
		}
	case "intent":
		if interaction.User.Intent == nil {
			return []interact.InteractionResponse{}, fmt.Errorf("intent request requires intent data")
		}
		// Convert test entities to interact entities
		var entities []interact.Entity
		for _, e := range interaction.User.Intent.Entities {
			entities = append(entities, interact.Entity{
				Name:  e.Name,
				Value: e.Value,
			})
		}
		interatctionRequest = interact.InteratctionRequest{
			Action: interact.Action{
				Type: "intent",
				Payload: interact.IntentPayload{
					Intent: interact.IntentData{
						Name: interaction.User.Intent.Name,
					},
					Entities: entities,
				},
			},
		}
	case "button":
		// Find the button by label from available buttons
		if len(availableButtons) == 0 {
			return []interact.InteractionResponse{}, fmt.Errorf("no buttons available from previous interaction")
		}
		var selectedButton *tests.Button
		for _, btn := range availableButtons {
			if btn.Request.Payload["label"] == interaction.User.Value {
				selectedButton = &btn
				break
			}
		}
		if selectedButton == nil {
			return []interact.InteractionResponse{}, fmt.Errorf("button with label '%s' not found in available buttons", interaction.User.Value)
		}
		// Send the button's request as the action
		interatctionRequest = interact.InteratctionRequest{
			Action: interact.Action{
				Type:    selectedButton.Request.Type,
				Payload: selectedButton.Request.Payload,
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
	// Pass the optional token to the interaction
	apiKey := global.VoiceflowAPIKey
	if apiKeyOverride != "" {
		apiKey = apiKeyOverride
	}
	req.Header.Add("Authorization", apiKey)
	req.Header.Add("versionID", environmentName)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error calling API: %v", err)
	}
	defer utils.SafeClose(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error reading response: %v", err)
	}

	var interactions []interact.InteractionResponse
	err = json.Unmarshal([]byte(string(body)), &interactions)
	if err != nil {
		return []interact.InteractionResponse{}, fmt.Errorf("error unmarshalling: %v. Response body: %s", err, string(body))
	}

	return interactions, nil
}
