package model

import (
	"net/http"
	"net/url"
	"strings"
)

type Proxies struct {
	Type  string   `json:"type"`
	URLs  []string `json:"urls,omitempty"`
}

const (
	ProxiesTypeNo     = "no"
	ProxiesTypeSystem = "system"
	ProxiesTypeList   = "list"
)

func NewProxies(proxyType string, urls []string) Proxies {
	return Proxies{
		Type: proxyType,
		URLs: filterValidProxies(urls),
	}
}

func (p *Proxies) FromString(s string) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", "no":
		p.Type = ProxiesTypeNo
	case "system":
		p.Type = ProxiesTypeSystem
	default:
		urls := strings.Split(s, ",")
		validURLs := filterValidProxies(urls)
		if len(validURLs) == 0 {
			p.Type = ProxiesTypeSystem
		} else {
			p.Type = ProxiesTypeList
			p.URLs = validURLs
		}
	}
}

func filterValidProxies(urls []string) []string {
	validURLs := make([]string, 0)
	for _, rawURL := range urls {
		rawURL = strings.TrimSpace(rawURL)
		if rawURL == "" {
			continue
		}
		// 尝试解析URL
		if _, err := url.Parse(rawURL); err == nil {
			validURLs = append(validURLs, rawURL)
		}
	}
	return validURLs
}

func (p *Proxies) GetClient() *http.Client {
	client := &http.Client{}
	switch p.Type {
	case ProxiesTypeNo:
		client.Transport = &http.Transport{
			Proxy: nil,
		}
	case ProxiesTypeSystem:
		// 使用系统默认代理
	case ProxiesTypeList:
		if len(p.URLs) > 0 {
			if proxyURL, err := url.Parse(p.URLs[0]); err == nil {
				client.Transport = &http.Transport{
					Proxy: http.ProxyURL(proxyURL),
				}
			}
		}
	}
	return client
}

func (p *Proxies) GetTransport() (*http.Transport, error) {
	if len(p.URLs) == 0 {
		return &http.Transport{}, nil
	}

	proxyURL, err := url.Parse(p.URLs[0])
	if err != nil {
		return nil, err
	}

	return &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}, nil
} 