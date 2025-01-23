package transcript

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/transcript"
	"github.com/xavidop/voiceflow-cli/pkg/voiceflow"
	"gopkg.in/yaml.v3"
)

func ToTest(agentID, transcriptID, outputFile, testName, testDescription string) error {

	transcriptJSON, err := voiceflow.FetchTranscriptJSON(agentID, transcriptID)
	if err != nil {
		return fmt.Errorf("failed to fetch transcript %s: %w", transcriptID, err)
	}

	testYaml, err := TranscriptToTest(transcriptJSON, testName, testDescription)
	if err != nil {
		return fmt.Errorf("failed to transform transcript %s to test: %w", transcriptID, err)
	}

	err = WriteYAMLToFile(testYaml, outputFile)
	if err != nil {
		return fmt.Errorf("failed to save transcript %s: %w", transcriptID, err)
	}

	return nil
}

func TranscriptToTest(transcriptJSON []transcript.Turn, testName, testDescription string) (tests.Test, error) {
	var test tests.Test = tests.Test{
		Name:         testName,
		Description:  testDescription,
		Interactions: []tests.Interaction{},
	}
	for index, turn := range transcriptJSON {
		if turn.Type == "launch" {
			agentResponse := findNextAgentTextResponse(transcriptJSON, index)
			test.Interactions = append(test.Interactions, tests.Interaction{
				ID: uuid.New().String(),
				User: tests.User{
					Text: "",
					Type: turn.Type,
				},
				Agent: tests.Agent{
					Validate: []tests.Validation{
						{
							ID:    uuid.New().String(),
							Type:  "equals",
							Value: agentResponse,
						},
					},
				},
			})
		} else if turn.Type == "request" {
			request, err := turn.Payload.GetIntentPayload()
			if err != nil {
				return tests.Test{}, fmt.Errorf("failed to get request payload message: %w", err)
			}
			agentResponse := findNextAgentTextResponse(transcriptJSON, index)
			test.Interactions = append(test.Interactions, tests.Interaction{
				ID: uuid.New().String(),
				User: tests.User{
					Text: request.Query,
					Type: "text",
				},
				Agent: tests.Agent{
					Validate: []tests.Validation{
						{
							ID:    uuid.New().String(),
							Type:  "equals",
							Value: agentResponse,
						},
					},
				},
			})
		}
	}
	return test, nil
}

func findNextAgentTextResponse(transcriptJSON []transcript.Turn, currentIndex int) string {
	for i := currentIndex + 1; i < len(transcriptJSON); i++ {
		if transcriptJSON[i].Type == "text" {
			response, err := transcriptJSON[i].Payload.GetTextPayload()
			if err == nil {
				return response.Message
			}
		}
	}
	return ""
}

func WriteYAMLToFile(data interface{}, filename string) error {
	// Convert to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data to YAML: %w", err)
	}

	// Write to file with 0644 permissions
	err = os.WriteFile(filename, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %w", err)
	}

	return nil
}
