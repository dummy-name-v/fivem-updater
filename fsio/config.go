package fsio

import (
	"encoding/json"
	"fivem-updater/github"
	"fmt"
	"log"
	"os"
)

type Config struct {
	filePath string
	Hash     string `json:"hash"`
}

func (c *Config) UpdateHash(hash string) error {
	c.Hash = hash

	buffer, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(c.filePath, buffer, os.ModePerm)
}

func GetConfig(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := &Config{filePath: filePath}
	err = json.Unmarshal(file, config)
	if err != nil {
		log.Println("unable to locate the updater config, creating it for you..")
		if err := os.WriteFile(filePath, []byte("{}"), os.ModePerm); err != nil {
			return nil, err
		}
	}

	return config, nil
}

func FormatConfigHash(platform *Platform, tag *github.Tag) string {
	return fmt.Sprintf("%s@%s", *platform, tag.Sha)
}
