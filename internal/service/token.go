package service

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/team-pascal/mnit/internal/model"
)

const TOKEN_ENV = "MNIT_NOTION_TOKEN"

func GetToken() (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", errors.New(err.Error())
	}

	token := config.Mnit.NotionToken
	if token == "" {
		return "", errors.New("not found mnit notion token.")
	}
	return token, nil
}

func SetToken(token string) error {
	file, err := GetConfigFile()
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()

	config := model.Config{
		Mnit: model.NotionConfig{NotionToken: token},
	}

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		return errors.New("failed to save token.")
	}
	return nil
}
