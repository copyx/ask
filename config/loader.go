package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	defaultConfigFileName = ".ask"
	envPrefix             = "ASK_"
)

var configPath = ""

func DefaultConfigDir() string {
	configDirPath, err := os.UserHomeDir()
	if err != nil {
		panic(newConfigError(err))
	}

	return configDirPath
}

func DefaultConfigPath() string {
	return filepath.Join(DefaultConfigDir(), defaultConfigFileName)
}

func ConfigDir() string {
	if configPath == "" {
		return DefaultConfigDir()
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		panic(newConfigError(err))
	}
	return filepath.Dir(absPath)
}

// SetConfigPath set path for configuration file
func SetConfigPath(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(newConfigError(err))
	} else if fileInfo.IsDir() {
		panic(newConfigError(fmt.Errorf("\"%s\" is not a file", path)))
	}

	configPath = path
}

func ConfigFileName() string {
	if configPath == "" {
		return defaultConfigFileName
	}

	absPath, err := filepath.Abs(configPath)
	if err != nil {
		panic(newConfigError(err))
	}
	return filepath.Base(absPath)
}

// InitConfig
func InitConfig() {
	viper.SetConfigName(ConfigFileName())
	viper.SetConfigType("toml")
	viper.AddConfigPath(ConfigDir())

	if err := viper.ReadInConfig(); err != nil {
		panic(newConfigError(err))
	}
}

func Get(key string) any {
	return viper.Get(key)
}

func newConfigError(err error) error {
	return fmt.Errorf("config: %w", err)
}
