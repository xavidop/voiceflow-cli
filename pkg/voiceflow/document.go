package voiceflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/document"
)

func UploadDocumentUrl(urlToUpload, name string, overwrite bool, maxChunkSize int, markdownConversion, llmGeneratedQ, llmPrependContext, llmBasedChunking, llmContentSummarization bool, tags []string) (string, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://api%s.voiceflow.com/v1/knowledge-base/docs/upload?", global.VoiceflowSubdomain)
	if overwrite {
		url = fmt.Sprintf("%soverwrite=true", url)
	}
	if maxChunkSize != 1000 {
		url = fmt.Sprintf("%s&maxChunkSize=%d", url, maxChunkSize)
	}
	if markdownConversion {
		url = fmt.Sprintf("%s&markdownConversion=true", url)
	}
	if llmGeneratedQ {
		url = fmt.Sprintf("%s&llmGeneratedQ=true", url)
	}
	if llmPrependContext {
		url = fmt.Sprintf("%s&llmPrependContext=true", url)
	}
	if llmBasedChunking {
		url = fmt.Sprintf("%s&llmBasedChunking=true", url)
	}
	if llmContentSummarization {
		url = fmt.Sprintf("%s&llmContentSummarization=true", url)
	}
	if len(tags) > 0 {
		url = fmt.Sprintf("%s&tags=[%s]", url, strings.Join(tags, ","))
	}

	analyticsRequest := document.URLDocument{
		Data: document.Data{
			Type: "url",
			Name: name,
			URL:  urlToUpload,
		},
	}

	byts, err := json.Marshal(analyticsRequest)
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

func UploadDocumentFile(fileToUpload string, overwrite bool, maxChunkSize int, markdownConversion, llmGeneratedQ, llmPrependContext, llmBasedChunking, llmContentSummarization bool, tags []string) (string, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	url := fmt.Sprintf("https://api%s.voiceflow.com/v1/knowledge-base/docs/upload?", global.VoiceflowSubdomain)
	if overwrite {
		url = fmt.Sprintf("%soverwrite=true", url)
	}
	if maxChunkSize != 1000 {
		url = fmt.Sprintf("%s&maxChunkSize=%d", url, maxChunkSize)
	}
	if markdownConversion {
		url = fmt.Sprintf("%s&markdownConversion=true", url)
	}
	if llmGeneratedQ {
		url = fmt.Sprintf("%s&llmGeneratedQ=true", url)
	}
	if llmPrependContext {
		url = fmt.Sprintf("%s&llmPrependContext=true", url)
	}
	if llmBasedChunking {
		url = fmt.Sprintf("%s&llmBasedChunking=true", url)
	}
	if llmContentSummarization {
		url = fmt.Sprintf("%s&llmContentSummarization=true", url)
	}
	if len(tags) > 0 {
		url = fmt.Sprintf("%s&tags=[%s]", url, strings.Join(tags, ","))
	}
	// Open the file
	fileContent, err := os.ReadFile(fileToUpload)
	if err != nil {
		return "", err
	}

	file, err := os.Open(fileToUpload)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a buffer to hold the multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileName := filepath.Base(fileToUpload)

	// Write the encoded content in the required format to the form field
	contentType := http.DetectContentType(fileContent)

	// Create a custom form field with filename in Content-Disposition
	partHeaders := make(textproto.MIMEHeader)
	partHeaders.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	partHeaders.Set("Content-Type", contentType)

	part, err := writer.CreatePart(partHeaders)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("error copying file content: %v", err)
	}

	// Close the multipart writer to set the boundary
	err = writer.Close()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", writer.FormDataContentType())
	req.Header.Add("Authorization", global.VoiceflowAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyResponse), nil
}

// ListDocuments fetches all documents in a knowledge base
func ListDocuments(page int, limit int, documentType string, includeTags []string, excludeTags []string, includeAllNonTagged bool, includeAllTagged bool) (string, error) {
	if global.VoiceflowSubdomain != "" {
		global.VoiceflowSubdomain = "." + global.VoiceflowSubdomain
	}
	baseURL := fmt.Sprintf("https://api%s.voiceflow.com/v1/knowledge-base/docs", global.VoiceflowSubdomain)

	// Build query parameters
	values := url.Values{}
	if page > 0 {
		values.Add("page", strconv.Itoa(page))
	}
	if limit > 0 {
		values.Add("limit", strconv.Itoa(limit))
	}
	if documentType != "" {
		values.Add("documentType", documentType)
	}
	if len(includeTags) > 0 {
		for _, tag := range includeTags {
				values.Add("includeTags", tag)
		}
	}
	if len(excludeTags) > 0 {
		for _, tag := range excludeTags {
			values.Add("excludeTags", tag)
		}
	}
	if includeAllNonTagged {
		values.Add("includeAllNonTagged", "true")
	}
	if includeAllTagged {
		values.Add("includeAllTagged", "true")
	}

	// Add query parameters to URL if any exist
	if len(values) > 0 {
		baseURL = baseURL + "?" + values.Encode()
	}

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
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
