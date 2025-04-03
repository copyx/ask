package config

import (
	"fmt"
	"os"
	"reflect"
)

type Configurations struct {
	GEMINI_API_KEY string
}

const envPrefix = "ASK_"

var configKeys = [...]string{"GEMINI_API_KEY"}

// loadConfigurations load configurations from env with validation
func (c *Configurations) LoadConfigurations() error {
	configValue := reflect.Indirect(reflect.ValueOf(c))

	for _, key := range configKeys {
		fieldValue := configValue.FieldByName(key)
		envName := envPrefix + key
		envValue := os.Getenv(envName)

		if envValue == "" {
			return fmt.Errorf("%+v is empty. Please set the env variable", envName)
		}

		fieldValue.SetString(envValue)
	}

	return nil
}
