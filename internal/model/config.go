package model

type Config struct {
	Mnit NotionConfig `toml:"mnit"`
}

type NotionConfig struct {
	NotionToken string            `toml:"notion_token"`
	NotionDBId  map[string]string `toml:"notion_db_id"`
}
