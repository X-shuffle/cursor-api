package application

import (
	"bytes"
	"cursor-api/internal/domain"
	"cursor-api/internal/domain/constants"
	"cursor-api/internal/domain/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ChatService struct {
	tokenRepo domain.TokenRepository
	client    *http.Client
	apiURL    string
}

func NewChatService(tokenRepo domain.TokenRepository) *ChatService {
	apiURL := constants.OPENAI_API_URL
	if constants.USE_REVERSE_PROXY {
		apiURL = "https://" + constants.REVERSE_PROXY_HOST + "/v1/chat/completions"
	}

	return &ChatService{
		tokenRepo: tokenRepo,
		client: &http.Client{
			Timeout: constants.SERVICE_TIMEOUT,
		},
		apiURL: apiURL,
	}
}

func (s *ChatService) HandleChat(req domain.ChatRequest) (interface{}, error) {
	// 获取可用的令牌
	tokens, err := s.tokenRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	// 选择一个可用的令牌
	var selectedToken *model.TokenInfo
	for i := range tokens {
		if tokens[i].Available {
			selectedToken = &tokens[i]
			break
		}
	}

	if selectedToken == nil {
		return nil, fmt.Errorf("no available tokens")
	}

	// 准备请求体
	requestBody := map[string]interface{}{
		"model":    req.Model,
		"messages": req.Messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// 创建请求
	request, err := http.NewRequest("POST", s.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+selectedToken.Token)

	// 发送请求
	startTime := time.Now()
	response, err := s.client.Do(request)
	if err != nil {
		selectedToken.Available = false
		s.updateToken(*selectedToken)
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	// 读取响应
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// 检查响应状态
	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusUnauthorized {
			selectedToken.Available = false
			s.updateToken(*selectedToken)
		}
		return nil, fmt.Errorf("API request failed with status %d: %s", response.StatusCode, string(body))
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// 更新令牌使用情况
	selectedToken.Used++
	if err := s.updateToken(*selectedToken); err != nil {
		// 只记录错误，不影响响应
		fmt.Printf("Failed to update token usage: %v\n", err)
	}

	// 构建响应
	result["token_id"] = selectedToken.Token
	result["timing"] = map[string]interface{}{
		"total": time.Since(startTime).Seconds(),
	}

	return result, nil
}

func (s *ChatService) updateToken(token model.TokenInfo) error {
	tokens, err := s.tokenRepo.GetAll()
	if err != nil {
		return err
	}

	for i, t := range tokens {
		if t.Token == token.Token {
			tokens[i] = token
			break
		}
	}

	return s.tokenRepo.SaveAll(tokens)
}

func (s *ChatService) GetModels() ([]string, error) {
	return []string{
		"gpt-4",
		"gpt-4-32k",
		"gpt-3.5-turbo",
		"gpt-3.5-turbo-16k",
		"gpt-4-turbo-preview",
		"gpt-4-0125-preview",
		"gpt-3.5-turbo-0125",
	}, nil
}

func (s *ChatService) GetAvailableTokens() ([]model.TokenInfo, error) {
	tokens, err := s.tokenRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// 过滤可用的令牌
	availableTokens := make([]model.TokenInfo, 0)
	for _, token := range tokens {
		// TODO: 实现令牌可用性检查逻辑
		availableTokens = append(availableTokens, token)
	}

	return availableTokens, nil
}

func (s *ChatService) UpdateTokenUsage(token model.TokenInfo) error {
	// TODO: 实现更新令牌使用情况的逻辑
	tokens, err := s.tokenRepo.GetAll()
	if err != nil {
		return err
	}

	// 更新令牌信息
	updatedTokens := make([]model.TokenInfo, len(tokens))
	copy(updatedTokens, tokens)
	for i, t := range updatedTokens {
		if t.Token == token.Token {
			updatedTokens[i] = token
			break
		}
	}

	return s.tokenRepo.SaveAll(updatedTokens)
}
