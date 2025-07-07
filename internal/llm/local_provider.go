package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocalLLMProvider struct {
	Model string
}

func NewLocalLLMProvider(model string) *LocalLLMProvider {
	return &LocalLLMProvider{Model: model}
}

func (l *LocalLLMProvider) GenerateDocumentation(ctx context.Context, prompt string) (string, error) {
	fmt.Println("Inside local gendocs")
	requestBody := map[string]string{
		"model":  l.Model,
		"prompt": prompt,
	}

	bodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:11434/api/generate", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var fullResponse string
	decoder := json.NewDecoder(resp.Body)
	for {
		var result struct {
			Response string `json:"response"`
		}
		if err := decoder.Decode(&result); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		fullResponse += result.Response
	}

	fmt.Printf("fullResponse is %s", fullResponse)
	return fullResponse, nil
}
