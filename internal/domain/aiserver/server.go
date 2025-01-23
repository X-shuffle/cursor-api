package aiserver

import "cursor-api/internal/domain/model"

type AIServer interface {
    Chat(req model.ChatRequest) (*model.ChatResponse, error)
    Stream(req model.ChatRequest, ch chan<- model.ChatResponse) error
} 