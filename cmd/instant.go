package cmd

import (
	"ask/ai"
	"ask/config"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(instantCmd)
}

var instantCmd = &cobra.Command{
	Use:   "instant",
	Short: "Send only one prompt to AI",
	RunE: func(cmd *cobra.Command, args []string) error {
		value := config.Get("gemini.api_key")
		apiKey, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid api key: %v", value)
		}

		var ai ai.GeminiAI
		ai.Init(apiKey)
		result, err := ai.GenerateContent(strings.Join(args, " "))
		if err != nil {
			return err
		}

		fmt.Println(result)
		return nil
	},
}
