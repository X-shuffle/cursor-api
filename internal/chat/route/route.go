package route

import (
	"cursor-api/internal/service"
	"github.com/gin-gonic/gin"
)

type ChatRouter struct {
	chatService *service.ChatService
}

func NewChatRouter(chatService *service.ChatService) *ChatRouter {
	return &ChatRouter{
		chatService: chatService,
	}
}

func (r *ChatRouter) RegisterRoutes(router *gin.Engine) {
	chat := router.Group("/v1/chat")
	{
		chat.POST("/completions", r.handleChat)
		chat.POST("/completions/stream", r.handleStreamChat)
	}
}
