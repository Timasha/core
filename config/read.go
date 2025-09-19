package config

import (
	"encoding/json"
	"os"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
)

var (
	configPath string
	validate   = validator.New(validator.WithRequiredStructEnabled())
)

func init() {
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.json"
	}
}

func ReadFromFile(cfg any) (err error) {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}

	err = defaults.Set(cfg)
	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		return err
	}

	err = validate.Struct(cfg)
	if err != nil {
		return err
	}

	return nil
}
