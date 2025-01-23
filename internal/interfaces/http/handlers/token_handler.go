package handlers

import (
	"cursor-api/internal/domain"
	"cursor-api/internal/domain/model"
	"cursor-api/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	tokenRepo domain.TokenRepository
}

func NewTokenHandler(tokenRepo domain.TokenRepository) *TokenHandler {
	return &TokenHandler{tokenRepo: tokenRepo}
}

// HandleTokensPage 处理令牌页面请求
func (h *TokenHandler) HandleTokensPage(c *gin.Context) {
	c.HTML(http.StatusOK, "tokens.html", nil)
}

// HandleGetTokens 处理获取令牌列表请求
func (h *TokenHandler) HandleGetTokens(c *gin.Context) {
	tokens, err := h.tokenRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "failed to get tokens",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tokens": tokens,
		"tokens_count": len(tokens),
	})
}

// HandleAddTokens 处理添加令牌请求
func (h *TokenHandler) HandleAddTokens(c *gin.Context) {
	var request []model.TokenAddRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request format",
		})
		return
	}

	// 验证请求数据
	if len(request) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "no tokens provided",
		})
		return
	}

	// 获取当前的令牌列表
	currentTokens, err := h.tokenRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "failed to get current tokens",
		})
		return
	}

	// 创建现有token的集合
	existingTokens := make(map[string]bool)
	for _, token := range currentTokens {
		existingTokens[token.Token] = true
	}

	// 处理新的tokens
	var newTokens []model.TokenInfo
	for _, req := range request {
		// 验证必要字段
		if req.Token == "" {
			continue
		}

		// 解析和验证令牌
		parsedToken := utils.ParseToken(req.Token)
		if !utils.ValidateToken(parsedToken) {
			continue
		}

		if !existingTokens[parsedToken] {
			tokenInfo := model.TokenInfo{
				Token: parsedToken,
				Checksum: func() string {
					if req.Checksum != nil {
						return utils.GenerateChecksumWithRepair(*req.Checksum)
					}
					return utils.GenerateChecksumWithDefault()
				}(),
				Available: true,
				Used: 0,
			}
			newTokens = append(newTokens, tokenInfo)
		}
	}

	// 如果有新tokens，添加它们
	if len(newTokens) > 0 {
		allTokens := append(currentTokens, newTokens...)
		if err := h.tokenRepo.SaveAll(allTokens); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"message": "failed to save tokens",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tokens_count": len(currentTokens) + len(newTokens),
		"message": "tokens have been added and reloaded",
	})
}

// HandleDeleteTokens 处理删除令牌请求
func (h *TokenHandler) HandleDeleteTokens(c *gin.Context) {
	var request model.TokensDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request format",
		})
		return
	}

	// 获取当前的令牌列表
	currentTokens, err := h.tokenRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "failed to get current tokens",
		})
		return
	}

	// 创建要删除的tokens的集合
	tokensToDelete := make(map[string]bool)
	for _, token := range request.Tokens {
		tokensToDelete[token] = true
	}

	// 过滤tokens
	var filteredTokens []model.TokenInfo
	var failedTokens []string
	for _, token := range currentTokens {
		if !tokensToDelete[token.Token] {
			filteredTokens = append(filteredTokens, token)
		} else {
			failedTokens = append(failedTokens, token.Token)
		}
	}

	// 保存更新后的令牌列表
	if err := h.tokenRepo.SaveAll(filteredTokens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "failed to save updated tokens",
		})
		return
	}

	response := gin.H{
		"status": "success",
	}

	// 根据期望添加额外信息
	if request.Expectation.NeedsUpdatedTokens() {
		var updatedTokens []string
		for _, token := range filteredTokens {
			updatedTokens = append(updatedTokens, token.Token)
		}
		response["updated_tokens"] = updatedTokens
	}
	if request.Expectation.NeedsFailedTokens() {
		response["failed_tokens"] = failedTokens
	}

	c.JSON(http.StatusOK, response)
}

// HandleReloadTokens 处理重新加载令牌请求
func (h *TokenHandler) HandleReloadTokens(c *gin.Context) {
	tokens, err := h.tokenRepo.Reload()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "failed to reload tokens",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tokens_count": len(tokens),
		"message": "token list has been reloaded",
	})
}

// HandleUpdateTokens 处理更新令牌请求
func (h *TokenHandler) HandleUpdateTokens(c *gin.Context) {
	var request struct {
		Tokens string `json:"tokens"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid request format",
		})
		return
	}

	// 解析令牌列表
	var tokens []model.TokenInfo
	if err := json.Unmarshal([]byte(request.Tokens), &tokens); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "invalid tokens format",
		})
		return
	}

	// 保存更新后的令牌列表
	if err := h.tokenRepo.SaveAll(tokens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": "failed to save tokens",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tokens_count": len(tokens),
		"message": "token files have been updated and reloaded",
	})
}
