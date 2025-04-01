package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/team-pascal/mnit/internal/infrastructure"
	"github.com/team-pascal/mnit/internal/model"
)

func Push(files []string) error {

	for _, path := range files {
		if !exists(path) {
			return fmt.Errorf("the file '%s' was not found.", path)
		}
		ex := filepath.Ext(path)
		if ex != ".md" {
			return errors.New("invalid file format. markdown files are allowed.")
		}
	}

	// Added process to convert files to notion api

	// sample
	data := model.Page{
		Parent: model.Parent{
			DatabaseId: "",
		},
		Properties: model.Properties{
			CompanyName: model.CompanyName{
				Title: []model.Title{
					{
						Text: model.Text{
							Content: "こんちは",
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to convert file. %w", err)
	}

	token, err := GetToken()
	if err != nil {
		return err
	}

	if err := infrastructure.Post(infrastructure.PagesEndpoint, bytes.NewBuffer(jsonData), token); err != nil {
		return err
	}

	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
