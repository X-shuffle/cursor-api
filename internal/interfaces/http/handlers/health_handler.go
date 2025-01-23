package handlers

import (
	"cursor-api/internal/application"
	"cursor-api/internal/domain/constants"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	appState    *application.AppState
	startTime   time.Time
	authToken   string
	pkgVersion  string
}

func NewHealthHandler(appState *application.AppState, authToken string) *HealthHandler {
	return &HealthHandler{
		appState:    appState,
		startTime:   time.Now(),
		authToken:   authToken,
		pkgVersion:  constants.PKG_VERSION,
	}
}

// HandleRoot 处理根路由请求
func (h *HealthHandler) HandleRoot(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/health")
}

// HandleHealth 处理健康检查请求
func (h *HealthHandler) HandleHealth(c *gin.Context) {
	// 检查认证令牌
	authHeader := c.GetHeader("Authorization")
	isAuthorized := false
	if authHeader != "" {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isAuthorized = token == h.authToken
	}

	// 计算运行时间
	uptime := int64(time.Since(h.startTime).Seconds())

	response := gin.H{
		"status":  "healthy",
		"version": h.pkgVersion,
		"uptime":  uptime,
		"models":  constants.AVAILABLE_MODELS,
		"endpoints": []string{
			"/v1/chat",
			"/v1/models",
			"/tokens",
			"/tokens/get",
			"/tokens/update",
			"/tokens/add",
			"/tokens/delete",
			"/logs",
			"/env-example",
			"/config",
			"/static",
			"/about",
			"/readme",
			"/api",
			"/hash",
			"/checksum",
			"/timestamp",
			"/calibration",
			"/user-info",
			"/build-key",
		},
	}

	// 如果是授权用户，添加系统统计信息
	if isAuthorized {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		stats := gin.H{
			"started":         h.startTime.Format(time.RFC3339),
			"total_requests": h.appState.GetTotalRequests(),
			"active_requests": h.appState.GetActiveRequests(),
			"system": gin.H{
				"memory": gin.H{
					"rss": memStats.Sys, // 物理内存使用量
				},
				"cpu": gin.H{
					"usage": getCPUUsage(), // 获取CPU使用率
				},
			},
		}
		response["stats"] = stats
	}

	c.JSON(http.StatusOK, response)
}

// getCPUUsage 获取CPU使用率
func getCPUUsage() float64 {
	// 这里可以使用 github.com/shirou/gopsutil 来获取更准确的 CPU 使用率
	// 简单实现返回一个模拟值
	return 0.0
} 