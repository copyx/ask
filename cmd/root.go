package cmd

import (
	"ask/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgPath = ""

func init() {
	cobra.OnInitialize(onInitialize)
	defaultCfgFilePath := config.GetDefaultCfgFilePath()

	rootCmd.PersistentFlags().StringVar(
		&cfgPath,
		"config",
		"",
		fmt.Sprintf("Config file path (default: %s)", defaultCfgFilePath))
}

func onInitialize() {
	config.SetConfigFilePath(cfgPath)
	config.InitConfig()
}

var rootCmd = &cobra.Command{
	Use:   "ask",
	Short: "A simple AI CLI client",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
