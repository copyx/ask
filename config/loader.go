package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultConfigFileName = ".ask"
	envPrefix             = "ASK_"
)

var (
	configFilePath = ""
	configKeys     = [...]string{"GEMINI_API_KEY"}
)

func GetDefaultCfgFilePath() string {
	configDirPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return configDirPath
}

func GetConfigFilePath() string {
	if configFilePath == "" {
		return GetDefaultCfgFilePath()
	}

	return configFilePath
}

func SetConfigFilePath(path string) {
	configFilePath = path
}

func InitConfig() {
	viper.SetConfigName(defaultConfigFileName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(GetConfigFilePath())

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Get(key string) any {
	return viper.Get(key)
}
