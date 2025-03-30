package service

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

func GetDBID(key string) (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", errors.New(err.Error())
	}
	dbID, ok := config.Mnit.NotionDBId[key]
	if !ok {
		return "", fmt.Errorf("db id for key '%s' not found.", key)
	}
	return dbID, nil
}

func SetDBID(key string, id string) error {
	config, err := GetConfig()
	if err != nil {
		return errors.New(err.Error())
	}

	file, err := GetConfigFile()
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()

	config.Mnit.NotionDBId[key] = id

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		return errors.New("failed to save database ID.")
	}
	return nil
}
