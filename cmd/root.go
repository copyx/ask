package cmd

import (
	"ask/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath = ""
	rootCmd    = &cobra.Command{
		Use:   "ask",
		Short: "A simple AI CLI client",
	}
)

func init() {
	cobra.OnInitialize(onInitialize)

	rootCmd.PersistentFlags().StringVar(
		&configPath,
		"config",
		config.DefaultConfigPath(),
		"Config file path")
}

func onInitialize() {
	config.SetConfigPath(configPath)
	config.InitConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
