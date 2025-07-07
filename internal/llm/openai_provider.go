package llm

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAiProvider struct {
	client *openai.Client
}

func NewOpenAiProvider() *OpenAiProvider {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY is not set")
	}
	return &OpenAiProvider{
		client: openai.NewClient(apiKey),
	}
}

func (c *OpenAiProvider) GenerateDocumentation(ctx context.Context, prompt string) (string, error) {

	log.Println("Inside openai provider gen docs")
	// log.Printf("prompt is %s", prompt)
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo, // or GPT3.5Turbo
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are an expert DevOps documentation generator.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
