package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"cursor-api/internal/domain/constants"
	"cursor-api/internal/utils"

	"github.com/google/uuid"
)

const (
	// HTTP 头常量
	secFetchDest     = "sec-fetch-dest"
	secFetchMode     = "sec-fetch-mode"
	secFetchSite     = "sec-fetch-site"
	secGpc           = "sec-gpc"
	priority         = "priority"
	proxyHost        = "x-co"

	// HTTP 头的值
	one            = "1"
	encodings      = "gzip,br"
	valueAccept    = "*/*"
	valueLanguage  = "zh-CN"
	empty          = "empty"
	cors           = "cors"
	noCache        = "no-cache"
	sameOrigin     = "same-origin"
	keepAlive      = "keep-alive"
	trailers       = "trailers"
	uEq4           = "u=4"
	uaWin          = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
)

var (
	httpClient     *http.Client
	httpClientOnce sync.Once
)

// GetHTTPClient 返回单例的 HTTP 客户端
func GetHTTPClient() *http.Client {
	httpClientOnce.Do(func() {
		httpClient = &http.Client{
			Timeout: constants.SERVICE_TIMEOUT,
		}
	})
	return httpClient
}

// BuildCursorClient 构建 Cursor API 请求
func BuildCursorClient(authToken, checksum string) (*http.Request, error) {
	traceID := uuid.New().String()

	req, err := http.NewRequest("POST", constants.CURSOR_API2_CHAT_URL, nil)
	if err != nil {
		return nil, err
	}

	// 设置基本头部
	req.Header.Set("Content-Type", "application/connect+proto")
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("connect-accept-encoding", encodings)
	req.Header.Set("connect-protocol-version", one)
	req.Header.Set("User-Agent", "connect-es/1.6.1")
	req.Header.Set("x-amzn-trace-id", fmt.Sprintf("Root=%s", traceID))
	req.Header.Set("x-client-key", utils.GenerateHash())
	req.Header.Set("x-cursor-checksum", checksum)
	req.Header.Set("x-cursor-client-version", "0.42.5")
	req.Header.Set("x-cursor-timezone", "Asia/Shanghai")
	req.Header.Set("x-ghost-mode", "true")
	req.Header.Set("x-request-id", traceID)
	req.Header.Set("Connection", keepAlive)
	req.Header.Set("Transfer-Encoding", "chunked")

	// 设置代理相关头部
	if constants.USE_REVERSE_PROXY {
		req.Header.Set("Host", constants.REVERSE_PROXY_HOST)
		req.Header.Set(proxyHost, constants.CURSOR_API2_HOST)
	} else {
		req.Header.Set("Host", constants.CURSOR_API2_HOST)
	}

	return req, nil
}

// BuildProfileClient 构建获取 Stripe 账户信息的请求
func BuildProfileClient(authToken string) (*http.Request, error) {
	req, err := http.NewRequest("GET", constants.CURSOR_API2_STRIPE_URL, nil)
	if err != nil {
		return nil, err
	}

	// 设置基本头部
	req.Header.Set("sec-ch-ua", "\"Not-A.Brand\";v=\"99\", \"Chromium\";v=\"124\"")
	req.Header.Set("x-ghost-mode", "true")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.42.5 Chrome/124.0.6367.243 Electron/30.4.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("Accept", valueAccept)
	req.Header.Set("Origin", "vscode-file://vscode-app")
	req.Header.Set(secFetchSite, "cross-site")
	req.Header.Set(secFetchMode, cors)
	req.Header.Set(secFetchDest, empty)
	req.Header.Set("Accept-Encoding", encodings)
	req.Header.Set("Accept-Language", valueLanguage)
	req.Header.Set(priority, "u=1, i")

	// 设置代理相关头部
	if constants.USE_REVERSE_PROXY {
		req.Header.Set("Host", constants.REVERSE_PROXY_HOST)
		req.Header.Set(proxyHost, constants.CURSOR_API2_HOST)
	} else {
		req.Header.Set("Host", constants.CURSOR_API2_HOST)
	}

	return req, nil
}

// BuildUsageClient 构建获取使用情况的请求
func BuildUsageClient(userID, authToken string) (*http.Request, error) {
	sessionToken := fmt.Sprintf("%s%%3A%%3A%s", userID, authToken)

	req, err := http.NewRequest("GET", constants.CURSOR_USAGE_API_URL, nil)
	if err != nil {
		return nil, err
	}

	// 设置基本头部
	req.Header.Set("User-Agent", uaWin)
	req.Header.Set("Accept", valueAccept)
	req.Header.Set("Accept-Language", valueLanguage)
	req.Header.Set("Accept-Encoding", encodings)
	req.Header.Set("Referer", "https://cursor.sh/settings")
	req.Header.Set("DNT", one)
	req.Header.Set(secGpc, one)
	req.Header.Set(secFetchDest, empty)
	req.Header.Set(secFetchMode, cors)
	req.Header.Set(secFetchSite, sameOrigin)
	req.Header.Set("Connection", keepAlive)
	req.Header.Set("Pragma", noCache)
	req.Header.Set("Cache-Control", noCache)
	req.Header.Set("TE", trailers)
	req.Header.Set(priority, uEq4)
	req.Header.Set("Cookie", fmt.Sprintf("WorkosCursorSessionToken=%s", sessionToken))

	// 设置查询参数
	q := req.URL.Query()
	q.Add("user", userID)
	req.URL.RawQuery = q.Encode()

	// 设置代理相关头部
	if constants.USE_REVERSE_PROXY {
		req.Header.Set("Host", constants.REVERSE_PROXY_HOST)
		req.Header.Set(proxyHost, constants.CURSOR_HOST)
	} else {
		req.Header.Set("Host", constants.CURSOR_HOST)
	}

	return req, nil
}

type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

type ClientOption func(*HTTPClient)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *HTTPClient) {
		c.client.Timeout = timeout
	}
}

func WithHeader(key, value string) ClientOption {
	return func(c *HTTPClient) {
		c.headers[key] = value
	}
}

func NewHTTPClient(baseURL string, opts ...ClientOption) *HTTPClient {
	c := &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *HTTPClient) Get(ctx context.Context, path string) ([]byte, error) {
	return c.Request(ctx, http.MethodGet, path, nil)
}

func (c *HTTPClient) Post(ctx context.Context, path string, body interface{}) ([]byte, error) {
	return c.Request(ctx, http.MethodPost, path, body)
}

func (c *HTTPClient) Request(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置默认 headers
	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(data))
	}

	return data, nil
}

func (c *HTTPClient) Stream(ctx context.Context, method, path string, body interface{}) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置流式传输的 headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(data))
	}

	return resp.Body, nil
} 