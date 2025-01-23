package app

import (
	"cursor-api/internal/domain/aiserver"
	"cursor-api/internal/domain/model"
	"cursor-api/internal/infrastructure/logger"
	"cursor-api/internal/middleware"
	"cursor-api/internal/service"
	"os"
	"sync"
)

type App struct {
	sync.RWMutex
	config       *Config
	chatService  *service.ChatService
	chatMiddleware *middleware.ChatMiddleware
	logger       *logger.Logger
	tokens       []model.TokenInfo
}

func NewApp(config *Config) (*App, error) {
	// 初始化日志
	log, err := logger.NewLogger(config.LogPath)
	if err != nil {
		return nil, err
	}

	// 初始化中间件
	chatMiddleware := middleware.NewChatMiddleware()

	// 初始化 AI 服务器
	server := aiserver.NewOpenAIServer(os.Getenv("OPENAI_API_KEY"))

	// 初始化聊天服务
	chatService := service.NewChatService(server)

	return &App{
		config:        config,
		chatService:   chatService,
		chatMiddleware: chatMiddleware,
		logger:        log,
	}, nil
}

func (a *App) GetConfig() *Config {
	return a.config
}

func (a *App) GetChatService() *service.ChatService {
	return a.chatService
}

func (a *App) GetChatMiddleware() *middleware.ChatMiddleware {
	return a.chatMiddleware
}

func (a *App) GetLogger() *logger.Logger {
	return a.logger
}

func (a *App) Shutdown() error {
	return a.logger.Close()
} 