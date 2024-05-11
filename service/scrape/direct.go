package scrape

import (
	"MyGo-scraper/common/utils"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-shiori/go-readability"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/html/charset"
)

func (s *Scrape) directScrape(ctx echo.Context, rawUrl string) (*ScrapeResult, error) {
	reqContext := ctx.Request().Context()
	req, err := http.NewRequestWithContext(reqContext, http.MethodGet, rawUrl, nil)
	if err != nil {
		return nil, err
	}

	// 构建请求头
	headers := map[string]string{
		"Accept":        "text/html;q=0.9, application/xhtml+xml;q=0.8",
		"Cache-Control": "no-cache",
		"Connection":    "keep-alive",
		"Pragma":        "no-cache",
		"User-Agent":    utils.RandomUserAgent(),
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 请求网页
	resp, err := scrapeClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 响应体(resp.Body)是一个流式的数据源。这意味着它会占用一个打开的网络连接和相关的资源，直到被显式地关闭

	// 解析resp
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request status: %v", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return nil, fmt.Errorf("invalid Content-Type: %s", contentType)
	}

	// 解析html的数据
	domain := ""
	u, err := url.Parse(rawUrl)
	if err != nil {
		domain = u.Scheme + "://" + u.Host
	}

	options := &md.Options{
		GetAbsoluteURL: func(selec *goquery.Selection, rawURL, _ string) string {
			if !s.rewiseDomain {
				return rawURL
			}
			// 如果是相对路径，拼接成绝对路径
			if strings.HasPrefix(rawURL, "/") {
				return domain + rawURL
			}
			return rawURL
		},
	}

	reader, err := charset.NewReader(resp.Body, contentType)
	if err != nil {
		return nil, err
	}

	// 使用 go-readability 解析
	if s.readability {
		article, err := readability.FromReader(reader, u)
		if err != nil {
			return nil, err
		}

		converter := md.NewConverter("", true, options)
		content, err := converter.ConvertString(article.Content)
		if err != nil {
			return nil, err
		}

		return &ScrapeResult{
			Url:     rawUrl,
			Title:   article.Title,
			Content: content,
		}, nil
	}

	// 使用 goquery 解析
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	converter := md.NewConverter("", true, options)
	content := converter.Convert(doc.Selection)

	title := ""
	if titleSelect := doc.Find("title"); titleSelect != nil {
		title = titleSelect.Text()
	}

	return &ScrapeResult{
		Url:     rawUrl,
		Title:   title,
		Content: content,
	}, nil
}
