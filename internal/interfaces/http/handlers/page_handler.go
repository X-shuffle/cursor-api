package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) HandleConfigPage(c *gin.Context) {
	c.HTML(http.StatusOK, "config.html", nil)
}

func (h *PageHandler) HandleConfigUpdate(c *gin.Context) {
	var config struct {
		// 配置字段
	}
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *PageHandler) HandleStatic(c *gin.Context) {
	filePath := c.Param("filepath")
	c.File(filepath.Join("static", filePath))
}

func (h *PageHandler) HandleAbout(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/readme")
}

func (h *PageHandler) HandleReadme(c *gin.Context) {
	c.HTML(http.StatusOK, "readme.html", nil)
}

func (h *PageHandler) HandleApiPage(c *gin.Context) {
	c.HTML(http.StatusOK, "api.html", nil)
} 