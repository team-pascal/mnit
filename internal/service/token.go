package service

import (
	"errors"
	"os"
)

const TOKEN_ENV = "MNIT_NOTION_TOKEN"

func GetToken() (string, error) {
	token := os.Getenv(TOKEN_ENV)
	if token == "" {
		return "", errors.New("not found mnit notion token")
	}
	return token, nil
}
