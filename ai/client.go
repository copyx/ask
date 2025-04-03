package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiAI struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func (ai *GeminiAI) Init(apiKey string) error {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("failed to initialize Gemini client:\n%w", err)
	}
	ai.client = client

	ai.model = client.GenerativeModel("gemini-2.0-flash")

	return nil
}

func (ai GeminiAI) GenerateContent(prompt string) (string, error) {
	resp, err := ai.model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content:\n%w", err)
	}

	return extractTextContent(resp), nil
}

func extractTextContent(resp *genai.GenerateContentResponse) string {
	textParts := []string{}
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if text, ok := part.(genai.Text); ok {
					textParts = append(textParts, string(text))
				}
			}
		}
	}

	return strings.Join(textParts, "")
}
