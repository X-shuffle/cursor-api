package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许某些路由不需要认证
		publicPaths := []string{"/", "/health", "/static"}
		path := c.Request.URL.Path
		
		// 检查是否是公开路径
		for _, publicPath := range publicPaths {
			if strings.HasPrefix(path, publicPath) {
				c.Next()
				return
			}
		}

		// 获取认证头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "invalid authentication token",
			})
			c.Abort()
			return
		}

		// 验证 Bearer token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != authToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "invalid authentication token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
