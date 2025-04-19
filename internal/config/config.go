package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL    string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (c *Config) save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return errors.New("could not locate base folder path")
	}
	os.WriteFile(path, data, os.ModePerm)

	return nil
}

func (c *Config) Read() error {
	path, err := getConfigFilePath()
	if err != nil {
		return errors.New("could not locate base folder path")
	}
	configJson, err := os.ReadFile(path)
	if err != nil {
		return errors.New("could not read from file")
	}
	json.Unmarshal(configJson, c)

	return nil
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	err := c.save()

	return err
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
