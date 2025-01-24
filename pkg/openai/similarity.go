package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/tests"
)

func OpenAICheckSimilarity(message string, s []string, similarityConfig tests.SimilarityConfig) (float64, error) {
	if similarityConfig.Provider != "openai" {
		return 0.0, fmt.Errorf("unsupported provider: %s", similarityConfig.Provider)
	}

	// OpenAI API endpoint
	apiURL := "https://api.openai.com/v1/chat/completions"

	// Prepare the prompt for similarity comparison
	prompt := fmt.Sprintf(
		"Compare the input message: \"%s\" to the following strings: %v. Calculate the average similarity score (0 to 1) across all comparisons. Return only the numeric score without any additional text or explanation.",
		message, s,
	)

	// Create the request payload
	payload := map[string]interface{}{
		"model":       similarityConfig.Model,
		"temperature": similarityConfig.Temperature,
		"top_p":       similarityConfig.TopP,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant that calculates similarity scores and returns only numeric results."},
			{"role": "user", "content": prompt},
		},
	}

	// Serialize the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0.0, fmt.Errorf("failed to serialize payload: %w", err)
	}

	// Make the HTTP POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0.0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+global.OpenAIAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0.0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return 0.0, fmt.Errorf("failed to parse response: %w", err)
	}

	// Extract similarity scores from the response
	choices := response["choices"].([]interface{})
	if len(choices) == 0 {
		return 0.0, fmt.Errorf("no choices returned in the response")
	}

	// Assume the assistant returns a list of similarity scores in the response text
	// Parse the similarity score from the response text
	responseText := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	var avgSimilarity float64
	if _, err := fmt.Sscanf(responseText, "%f", &avgSimilarity); err != nil {
		return 0, fmt.Errorf("failed to parse similarity score: %w", err)
	}

	return avgSimilarity, nil

}
