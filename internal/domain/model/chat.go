package model

// Message 定义聊天消息
type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

// ChatRequest 定义聊天请求
type ChatRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}

// ChatResponse 定义聊天响应
type ChatResponse struct {
    ID      string `json:"id"`
    Object  string `json:"object"`
    Created int64  `json:"created"`
    Model   string `json:"model"`
    Usage   struct {
        PromptTokens     int `json:"prompt_tokens"`
        CompletionTokens int `json:"completion_tokens"`
        TotalTokens     int `json:"total_tokens"`
    } `json:"usage"`
    Choices []Choice `json:"choices"`
}

type Choice struct {
    Message      Message `json:"message"`
    FinishReason string  `json:"finish_reason"`
} 