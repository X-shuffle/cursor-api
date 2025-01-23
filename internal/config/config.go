package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	MaxTokens      int     `json:"max_tokens"`
	Temperature    float64 `json:"temperature"`
	BaseURL        string  `json:"base_url"`
	RequestTimeout int     `json:"request_timeout"`
}

var (
	config     *Config
	configOnce sync.Once
	configLock sync.RWMutex
)

func GetConfig() *Config {
	configOnce.Do(func() {
		config = &Config{
			MaxTokens:      2000,
			Temperature:    0.7,
			BaseURL:        "https://api.openai.com/v1",
			RequestTimeout: 30,
		}
		loadConfig()
	})
	return config
}

func loadConfig() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return
	}

	configLock.Lock()
	defer configLock.Unlock()

	if err := json.Unmarshal(data, config); err != nil {
		log.Printf("Failed to parse config: %v", err)
	}
}

func SaveConfig(cfg Config) error {
	configLock.Lock()
	defer configLock.Unlock()

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile("config.json", data, 0644); err != nil {
		return err
	}

	*config = cfg
	return nil
}
