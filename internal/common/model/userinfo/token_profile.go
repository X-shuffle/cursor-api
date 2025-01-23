package userinfo

type TokenProfile struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Permissions struct {
		CanUseGPT4    bool `json:"can_use_gpt4"`
		CanUseGPT4V   bool `json:"can_use_gpt4v"`
		CanUseClaudeV2 bool `json:"can_use_claude_v2"`
	} `json:"permissions"`
	Usage struct {
		GPT4Used     int64 `json:"gpt4_used"`
		GPT4Limit    int64 `json:"gpt4_limit"`
		ClaudeV2Used int64 `json:"claude_v2_used"`
		ClaudeV2Limit int64 `json:"claude_v2_limit"`
	} `json:"usage"`
} 