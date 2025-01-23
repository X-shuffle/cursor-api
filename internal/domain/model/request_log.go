package model

import "time"

type RequestLog struct {
	ID        int64     `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Model     string    `json:"model"`
	TokenInfo TokenInfo `json:"token_info"`
	Prompt    string    `json:"prompt,omitempty"`
	Timing    struct {
		Total float64 `json:"total"`
		First float64 `json:"first,omitempty"`
	} `json:"timing"`
	Stream bool   `json:"stream"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
} 