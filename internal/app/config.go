package app

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	MaxTokens      int     `json:"max_tokens"`
	Temperature    float64 `json:"temperature"`
	BaseURL        string  `json:"base_url"`
	RequestTimeout int     `json:"request_timeout"`
	TokenPath      string  `json:"token_path"`
	LogPath        string  `json:"log_path"`
	RateLimit      int     `json:"rate_limit"`
	RateLimitBurst int     `json:"rate_limit_burst"`
}

var (
	config     *Config
	configOnce sync.Once
	configLock sync.RWMutex
)

func LoadConfig(path string) (*Config, error) {
	configOnce.Do(func() {
		config = &Config{
			MaxTokens:      2000,
			Temperature:    0.7,
			BaseURL:        "https://api.openai.com/v1",
			RequestTimeout: 30,
			TokenPath:      "tokens.json",
			LogPath:        "logs/app.log",
			RateLimit:      10,
			RateLimitBurst: 30,
		}

		if path != "" {
			data, err := os.ReadFile(path)
			if err == nil {
				json.Unmarshal(data, config)
			}
		}
	})

	return config, nil
} 