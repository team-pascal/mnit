package main

import (
	"fmt"
	"os"

	"github.com/team-pascal/mnit/internal/service"
)

func main() {
	token, err := service.GetToken()
	if err != nil {
		printError(err.Error())
	}

	fmt.Println(token)
}

func printError(message string) {
	fmt.Fprintf(os.Stderr, "mnit: %v", message)
}
