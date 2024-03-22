package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

const CONFIG_FILE = "config.json"

type Config struct {
	ApiKey     string `json:"apiKey"`
	MerchantId string `json:"merchantId"`
}

func (c *Config) Validate() error {
	value := reflect.ValueOf(*c)

	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Interface() == "" {
			arg := value.Type().Field(i).Tag.Get("json")

			return errors.New(
				fmt.Sprintf("\"%s\" config field is empty", arg),
			)
		}
	}

	return nil
}

func Parse() (*Config, error) {
	if !IsExists() {
		err := GenerateInitialConfig()

		if err != nil {
			return nil, err
		}

		return nil, errors.New(
			"config file (config.json) not found. initial config file is created, fill it",
		)
	}

	bytes, _ := os.ReadFile(CONFIG_FILE)

	cfg := Config{}
	err := json.Unmarshal(bytes, &cfg)

	if err != nil {
		return nil, err
	}

	err = cfg.Validate()

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func IsExists() bool {
	_, err := os.Stat(CONFIG_FILE)

	return !os.IsNotExist(err)
}

func GenerateInitialConfig() error {
	file, err := os.Create(CONFIG_FILE)

	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	initialConfig, _ := json.MarshalIndent(Config{}, "", "\t")

	_, err = file.Write(initialConfig)

	if err != nil {
		return err
	}

	return nil
}
