package handlers

import (
	"cursor-api/internal/application"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	appState *application.AppState
}

func NewLogHandler(appState *application.AppState) *LogHandler {
	return &LogHandler{appState: appState}
}

// HandleLogs 处理日志页面请求
func (h *LogHandler) HandleLogs(c *gin.Context) {
	c.HTML(http.StatusOK, "logs.html", nil)
}

// HandleLogsPost 处理日志数据请求
func (h *LogHandler) HandleLogsPost(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"message": "unauthorized",
		})
		return
	}

	logs := h.appState.GetLogs()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"total": len(logs),
		"active": h.appState.GetActiveRequests(),
		"error": h.appState.GetErrorRequests(),
		"logs": logs,
		"timestamp": time.Now().Format(time.RFC3339),
	})
} 