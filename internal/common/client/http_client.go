package client

import (
	"context"
	"net/http"
	"time"
)

// RebuildHTTPClient 重建HTTP客户端
func RebuildHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
}

// RebuildHTTPClientWithContext 重建带上下文的HTTP客户端
func RebuildHTTPClientWithContext(ctx context.Context, timeout time.Duration) *http.Client {
	client := RebuildHTTPClient(timeout)
	
	// 使用上下文包装客户端的传输层
	if transport, ok := client.Transport.(*http.Transport); ok {
		transport.DialContext = (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext
	}

	return client
} 