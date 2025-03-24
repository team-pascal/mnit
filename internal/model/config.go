package model

type Config struct {
	Notion NotionConfig `toml:"notion"`
}

type NotionConfig struct {
	NotionToken string `toml:"notion_token"`
}
