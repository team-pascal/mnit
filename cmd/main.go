package main

import (
	"fmt"

	"github.com/team-pascal/mnit/internal/service"
)

func main() {
	token := service.GetToken()
	fmt.Println(token)
}
