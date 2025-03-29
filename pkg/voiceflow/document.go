package voiceflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/xavidop/voiceflow-cli/internal/global"
	"github.com/xavidop/voiceflow-cli/internal/types/voiceflow/document"
	"github.com/xavidop/voiceflow-cli/internal/utils"
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
	defer utils.SafeClose(res.Body)

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
	defer utils.SafeClose(file)

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
	defer utils.SafeClose(res.Body)

	bodyResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyResponse), nil
}
