package service

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/team-pascal/mnit/internal/model"
)

func GetConfig() (model.Config, error) {
	configFilePath, err := getConfigPath()
	if err != nil {
		return model.Config{}, errors.New(err.Error())
	}

	var config model.Config

	if _, err := toml.DecodeFile(configFilePath, &config); err != nil {
		return model.Config{}, errors.New("failed open config file.")
	}

	return config, nil
}

func GetConfigFile() (*os.File, error) {
	configFilePath, err := getConfigPath()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	file, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, errors.New("failed open config file.")
	}

	return file, nil
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New(err.Error())
	}
	configDir := filepath.Join(home, ".mnit")
	configFile := filepath.Join(configDir, "config.toml")

	return configFile, nil
}
