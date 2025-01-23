package model

// PageContent 页面内容类型
type PageContent struct {
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
}

const (
	PageContentTypeDefault = "default"
	PageContentTypeText   = "text"
	PageContentTypeHTML   = "html"
)

// NewDefaultPageContent 创建默认页面内容
func NewDefaultPageContent() PageContent {
	return PageContent{Type: PageContentTypeDefault}
}

// NewTextPageContent 创建文本页面内容
func NewTextPageContent(content string) PageContent {
	return PageContent{
		Type:    PageContentTypeText,
		Content: content,
	}
}

// NewHTMLPageContent 创建HTML页面内容
func NewHTMLPageContent(content string) PageContent {
	return PageContent{
		Type:    PageContentTypeHTML,
		Content: content,
	}
} 