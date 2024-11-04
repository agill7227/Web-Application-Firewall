package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DetectionConfig struct {
	BlockedUserAgentsFile string `yaml:"blocked_user_agents_file"`
	BlockedPathsFile      string `yaml:"blocked_paths_file"`
	XssPayloadsFile       string `yaml:"xss_payloads_file"`
}

type RateLimitsConfig struct {
}

type Config struct {
}

var AppConfig Config

func LoadConfig(configFile string) Config {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Failed to decode config file: %s", err)
	}

	return AppConfig

}
