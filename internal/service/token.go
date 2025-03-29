package service

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/team-pascal/mnit/internal/model"
)

const TOKEN_ENV = "MNIT_NOTION_TOKEN"

func GetToken() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New(err.Error())
	}
	configDir := filepath.Join(home, ".mnit")
	configFile := filepath.Join(configDir, "config.toml")

	var config model.Config

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return "", errors.New("failed open config file.")
	}

	token := config.Notion.NotionToken
	if token == "" {
		return "", errors.New("not found mnit notion token.")
	}
	return token, nil
}

func SetToken(token string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return errors.New(err.Error())
	}
	configDir := filepath.Join(home, ".mnit")
	configFile := filepath.Join(configDir, "config.toml")
	file, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("failed open config file.")
	}
	defer file.Close()

	config := model.Config{
		Notion: model.NotionConfig{NotionToken: token},
	}

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		return errors.New("failed to save token.")
	}
	return nil
}
