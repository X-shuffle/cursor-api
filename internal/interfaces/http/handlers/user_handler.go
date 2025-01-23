package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// HandleUserInfo 处理用户信息请求
func (h *UserHandler) HandleUserInfo(c *gin.Context) {
	var req struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request",
		})
		return
	}

	// 获取用户信息的逻辑...
} 