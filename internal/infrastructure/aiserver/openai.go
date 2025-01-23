package aiserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"cursor-api/internal/domain/model"
	"cursor-api/internal/infrastructure/client"
	"bufio"
	"strings"
)

type OpenAIServer struct {
	client  *client.HTTPClient
}

func NewOpenAIServer(apiKey string) *OpenAIServer {
	httpClient := client.NewHTTPClient(
		"https://api.openai.com/v1",
		client.WithHeader("Authorization", "Bearer "+apiKey),
	)
	return &OpenAIServer{
		client: httpClient,
	}
}

func (s *OpenAIServer) Chat(req model.ChatRequest) (*model.ChatResponse, error) {
	data, err := s.client.Post(context.Background(), "/chat/completions", req)
	if err != nil {
		return nil, err
	}

	var response model.ChatResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *OpenAIServer) Stream(req model.ChatRequest, ch chan<- model.ChatResponse) error {
	// 添加流式处理标志
	streamReq := struct {
		model.ChatRequest
		Stream bool `json:"stream"`
	}{
		ChatRequest: req,
		Stream:      true,
	}

	// 获取流式响应
	reader, err := s.client.Stream(context.Background(), http.MethodPost, "/chat/completions", streamReq)
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "data: [DONE]" {
			continue
		}

		// 移除 "data: " 前缀
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		line = strings.TrimPrefix(line, "data: ")

		var response model.ChatResponse
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			return fmt.Errorf("failed to parse SSE: %w", err)
		}

		ch <- response
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("stream read failed: %w", err)
	}

	return nil
} 