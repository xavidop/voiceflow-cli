package voiceflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
	kbTypes "github.com/xavidop/voiceflow-cli/internal/types/voiceflow/kb"
)

func QueryKB(question, model string, temperature float64, chunkLimit int, synthesis bool, systemPrompt string, includeTags []string, includeOperator string, excludeTags []string, excludeOperator string, includeAllTagged bool, includeAllNonTagged bool) (string, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://general-runtime%s.voiceflow.com/knowledge-base/query", global.VoiceflowSubdomain)

	if includeOperator != "" && includeOperator != "and" && includeOperator != "or" && excludeOperator != "" && excludeOperator != "and" && excludeOperator != "or" {
		return "", fmt.Errorf("invalid operator, must be and/or")
	}

	var tags *kbTypes.Tags
	if includeAllTagged {
		tags = &kbTypes.Tags{
			IncludeAllTagged: true,
		}
	}
	if includeAllNonTagged {
		if tags == nil {
			tags = &kbTypes.Tags{
				IncludeAllNonTagged: true,
			}
		} else {
			tags.IncludeAllNonTagged = true
		}
	}
	if len(includeTags) > 0 {
		if tags == nil {
			tags = &kbTypes.Tags{
				Include: &kbTypes.TagOperator{
					Items:    includeTags,
					Operator: includeOperator,
				},
			}
		} else {
			tags.Include = &kbTypes.TagOperator{
				Items:    includeTags,
				Operator: includeOperator,
			}
		}
	}

	if len(excludeTags) > 0 {
		if tags == nil {
			tags = &kbTypes.Tags{
				Exclude: &kbTypes.TagOperator{
					Items:    excludeTags,
					Operator: excludeOperator,
				},
			}
		} else {
			tags.Exclude = &kbTypes.TagOperator{
				Items:    excludeTags,
				Operator: excludeOperator,
			}
		}
	}

	kbRequest := kbTypes.Query{
		ChunkLimit: chunkLimit,
		Question:   question,
		Synthesis:  synthesis,
		Tags:       tags,
		Settings: &kbTypes.Settings{
			Model:       model,
			Temperature: temperature,
			System:      systemPrompt,
		},
	}

	byts, err := json.Marshal(kbRequest)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %v", err)
	}

	payload := strings.NewReader(string(byts))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", global.VoiceflowAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SaveQuery(analytics string, outputFile string) error {

	// Write agent to file
	err := os.WriteFile(outputFile, []byte(analytics), 0644)
	if err != nil {
		return fmt.Errorf("failed to write analytcs file: %w", err)
	}

	return nil
}
