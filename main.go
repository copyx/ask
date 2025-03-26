package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	if len(os.Args) < 2 {
		printUsages()
		os.Exit(1)
	}

	var config Configurations
	err := config.loadConfigurations()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args := os.Args[1:]
	question := strings.Join(args, " ")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GEMINI_API_KEY))
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

type Configurations struct {
	GEMINI_API_KEY string
}

const envPrefix = "ASK_"

var configKeys = [...]string{"GEMINI_API_KEY"}

// loadConfigurations load configurations from env with validation
func (c *Configurations) loadConfigurations() error {
	configValue := reflect.Indirect(reflect.ValueOf(c))

	for _, key := range configKeys {
		fieldValue := configValue.FieldByName(key)
		envName := envPrefix + key
		envValue := os.Getenv(envName)

		if envValue == "" {
			return errors.New(fmt.Sprintf("%+v is empty. Please set the env variable.", envName))
		}

		fieldValue.SetString(envValue)
	}

	return nil
}
