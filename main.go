package main

import (
	"cursor-api/internal/application"
	"cursor-api/internal/domain/constants"
	"cursor-api/internal/infrastructure/logger"
	"cursor-api/internal/infrastructure/persistence"
	"cursor-api/internal/interfaces/http/handlers"
	"cursor-api/internal/interfaces/http/middleware"
	"cursor-api/internal/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

var authToken string

func init() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 初始化 auth token
	authToken = os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("AUTH_TOKEN must be set")
	}
}

func main() {
	// 初始化应用状态
	tokens, err := utils.LoadTokens()
	if err != nil {
		log.Fatal("Failed to load tokens:", err)
	}
	appState := application.NewAppState(tokens)

	// 初始化仓储
	tokenRepo := persistence.NewFileTokenRepository("tokens.json")

	// 初始化服务
	chatService := application.NewChatService(tokenRepo)

	// 初始化处理器
	healthHandler := handlers.NewHealthHandler(appState, authToken)
	chatHandler := handlers.NewChatHandler(chatService)
	miscHandler := handlers.NewMiscHandler(appState)
	tokenHandler := handlers.NewTokenHandler(tokenRepo)
	pageHandler := handlers.NewPageHandler()

	// 设置路由
	r := gin.Default()

	// CORS 配置
	r.Use(cors.Default())

	// 请求体大小限制
	maxBodySize := getEnvInt("REQUEST_BODY_LIMIT_MB", 2)
	r.MaxMultipartMemory = int64(maxBodySize) << 20 // maxBodySize MB

	// 初始化日志
	logger, err := logger.NewLogger("logs/app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()

	// 初始化速率限制器
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 30, 1*time.Hour)

	// 添加中间件
	r.Use(middleware.AuthMiddleware(authToken))
	r.Use(rateLimiter.RateLimit())

	// 记录请求日志
	r.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		logger.Log("INFO", "Request completed",
			map[string]interface{}{
				"path":     path,
				"method":   c.Request.Method,
				"status":   c.Writer.Status(),
				"duration": time.Since(start),
				"ip":       c.ClientIP(),
			},
		)
	})

	// 静态文件和模板配置
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// 路由配置
	r.GET("/", healthHandler.HandleRoot)
	r.GET("/health", healthHandler.HandleHealth)
	r.GET("/tokens", tokenHandler.HandleTokensPage)
	r.GET("/models", chatHandler.HandleModels)
	r.POST("/tokens/get", tokenHandler.HandleGetTokens)
	r.POST("/tokens/reload", tokenHandler.HandleReloadTokens)
	r.POST("/tokens/update", tokenHandler.HandleUpdateTokens)
	r.POST("/tokens/add", tokenHandler.HandleAddTokens)
	r.POST("/tokens/delete", tokenHandler.HandleDeleteTokens)
	r.POST("/chat", chatHandler.HandleChat)
	r.GET(constants.ROUTE_LOGS_PATH, miscHandler.HandleLogs)
	r.POST(constants.ROUTE_LOGS_PATH, miscHandler.HandleLogsPost)
	r.GET(constants.ROUTE_ENV_EXAMPLE_PATH, miscHandler.HandleEnvExample)
	r.GET(constants.ROUTE_GET_HASH, miscHandler.HandleGetHash)
	r.GET(constants.ROUTE_GET_CHECKSUM, miscHandler.HandleGetChecksum)
	r.GET(constants.ROUTE_GET_TIMESTAMP_HEADER, miscHandler.HandleGetTimestampHeader)
	r.POST(constants.ROUTE_BASIC_CALIBRATION_PATH, miscHandler.HandleBasicCalibration)
	r.POST(constants.ROUTE_USER_INFO_PATH, miscHandler.HandleUserInfo)
	r.GET(constants.ROUTE_BUILD_KEY_PATH, miscHandler.HandleBuildKeyPage)
	r.POST(constants.ROUTE_BUILD_KEY_PATH, miscHandler.HandleBuildKey)
	r.GET(constants.ROUTE_CONFIG_PATH, pageHandler.HandleConfigPage)
	r.POST(constants.ROUTE_CONFIG_PATH, pageHandler.HandleConfigUpdate)
	r.GET(constants.ROUTE_ABOUT_PATH, pageHandler.HandleAbout)
	r.GET(constants.ROUTE_README_PATH, pageHandler.HandleReadme)
	r.GET(constants.ROUTE_API_PATH, pageHandler.HandleApiPage)
	r.GET(constants.ROUTE_MODELS_PATH, chatHandler.HandleModels)
	r.POST(constants.ROUTE_CHAT_PATH, chatHandler.HandleChatCompletions)

	// 启动服务器
	port := getEnvString("PORT", "3000")
	addr := fmt.Sprintf("0.0.0.0:%s", port)

	fmt.Printf("Server running on port %s\n", port)
	fmt.Printf("Current version: v%s\n", constants.PKG_VERSION)
	fmt.Println("This is a beta version, please report any issues!")

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

// 工具函数
func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}
