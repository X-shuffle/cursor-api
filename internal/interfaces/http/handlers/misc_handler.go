package handlers

import (
	"bufio"
	"crypto/md5"
	"cursor-api/internal/application"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type MiscHandler struct {
	appState *application.AppState
	logPath  string
}

func NewMiscHandler(appState *application.AppState) *MiscHandler {
	return &MiscHandler{
		appState: appState,
		logPath:  "logs/app.log", // 可以通过配置注入
	}
}

func (h *MiscHandler) HandleLogs(c *gin.Context) {
	// 读取日志文件
	file, err := os.Open(h.logPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open log file"})
		return
	}
	defer file.Close()

	// 读取最后的 N 行日志
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	maxLines := 100 // 可配置

	// 读取所有行
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > maxLines {
			lines = lines[1:] // 保持最大行数
		}
	}

	if err := scanner.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read logs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": lines})
}

func (h *MiscHandler) HandleLogsPost(c *gin.Context) {
	// 获取日志内容
	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Dir(h.logPath), 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create log directory"})
		return
	}

	// 追加日志
	file, err := os.OpenFile(h.logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open log file"})
		return
	}
	defer file.Close()

	timestamp := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("[%s] USER: %s\n", timestamp, req.Content)

	if _, err := file.WriteString(logEntry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *MiscHandler) HandleEnvExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"AUTH_TOKEN":            "your_auth_token",
		"PORT":                  "3000",
		"REQUEST_BODY_LIMIT_MB": "2",
		"TOKEN_PATH":            "tokens.json",
		"LOG_PATH":              "logs/app.log",
		"RATE_LIMIT":            "10",
		"RATE_LIMIT_BURST":      "30",
	})
}

func (h *MiscHandler) HandleGetHash(c *gin.Context) {
	data := fmt.Sprintf("%d", time.Now().Unix())
	hash := md5.Sum([]byte(data))
	c.String(http.StatusOK, fmt.Sprintf("%x", hash))
}

func (h *MiscHandler) HandleGetChecksum(c *gin.Context) {
	// 生成新的校验和
	timestamp := time.Now().Unix()
	checksum := generateChecksum(timestamp)
	c.String(http.StatusOK, checksum)
}

func (h *MiscHandler) HandleGetTimestampHeader(c *gin.Context) {
	c.Header("X-Timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	c.Status(http.StatusOK)
}

func (h *MiscHandler) HandleBasicCalibration(c *gin.Context) {
	var req struct {
		Model string `json:"model"`
		Text  string `json:"text"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 执行基础校准
	result := map[string]interface{}{
		"status":     "calibrated",
		"model":      req.Model,
		"text":       req.Text,
		"timestamp":  time.Now().Unix(),
		"calibrated": true,
	}

	c.JSON(http.StatusOK, result)
}

func (h *MiscHandler) HandleUserInfo(c *gin.Context) {
	var req struct {
		UserID string                 `json:"user_id"`
		Info   map[string]interface{} `json:"info"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 处理用户信息
	response := map[string]interface{}{
		"user_id":   req.UserID,
		"info":      req.Info,
		"processed": true,
		"timestamp": time.Now().Unix(),
	}

	c.JSON(http.StatusOK, response)
}

func (h *MiscHandler) HandleBuildKeyPage(c *gin.Context) {
	c.HTML(http.StatusOK, "build_key.html", gin.H{
		"title": "Build API Key",
	})
}

func (h *MiscHandler) HandleBuildKey(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成 API Key
	timestamp := time.Now().UnixNano()
	data := fmt.Sprintf("%s-%s-%d", req.Prefix, req.Name, timestamp)
	hash := md5.Sum([]byte(data))
	apiKey := fmt.Sprintf("%s_%x", req.Prefix, hash)

	c.JSON(http.StatusOK, gin.H{
		"key":     apiKey,
		"name":    req.Name,
		"created": time.Now().Unix(),
	})
}

// generateChecksum 生成校验和
func generateChecksum(timestamp int64) string {
	// 实现校验和生成逻辑
	return time.Unix(timestamp, 0).Format("20060102150405")
}
