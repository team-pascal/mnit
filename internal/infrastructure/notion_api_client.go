package infrastructure

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const NOTION_API_URL = "https://api.notion.com/v1/"
const NOTION_VERSION = "2022-06-28"
const mimeType = "application/json"

const (
	PagesEndpoint = NOTION_API_URL + "pages"
)

func Post(url string, body io.Reader, token string) error {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return errors.New("failed to create request.")
	}
	req.Header.Set("Content-Type", mimeType)
	req.Header.Set("Notion-Version", NOTION_VERSION)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to post data: %w", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("faile to read response body. %w", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("failed to post. status: %d, response: %s", res.StatusCode, string(resBody))
	}

	return err
}
