package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	if len(os.Args) < 2 {
		printUsages()
		os.Exit(1)
	}

	args := os.Args[1:]
	question := strings.Join(args, " ")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func printUsages() {
	usages := []string{
		"usages: ask <prompt>",
		"",
		"This is an AI client for CLI.",
		"You need to set GEMINI_API_KEY environment variable with an API key for Gemini.",
	}

	for _, s := range usages {
		fmt.Println(s)
	}
}
