package service

import (
	"fmt"
	"os"
)

const TOKEN_ENV = "MNIT_NOTION_TOKEN"

func GetToken() string {
	token := os.Getenv(TOKEN_ENV)
	if token == "" {
		fmt.Fprintf(os.Stderr, "mnit: not found mnit notion token")
	}
	return token
}
