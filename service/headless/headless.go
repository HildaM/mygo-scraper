package headless

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

// Headless 函数接收网址和两个布尔参数，返回网页的标题、内容和可能的错误。
func Headless(rawUrl string, readability, rewiseDomain bool) (string, string, error) {
	ctx, cancel := chromedp.NewContext(allocCtx) // 创建一个新的chromedp上下文
	defer cancel()                               // 函数结束时取消上下文

	// 初始化chromedp运行环境
	_ = chromedp.Run(ctx)

	// 设置超时上下文，限制操作时间为10秒
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 定义用于存储执行JavaScript脚本结果的变量
	var result map[string]any

	// 默认的JavaScript脚本，获取页面标题和HTML内容
	script := "(function __webscraper__(){return {title: document.title, content: document.documentElement.outerHTML};})();"
	if readability {
		// 如果启用readability，则使用Readability.js解析网页
		script = readabilityJS + ";new Readability(document.cloneNode(true)).parse();"
	}

	// 使用chromedp执行导航和脚本评估, 获取rawUrl网页的html
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(rawUrl), // 导航到指定的URL
		chromedp.Evaluate(
			script,  // 执行JavaScript脚本
			&result, // 将结果存储到result变量中
			func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
				return p.WithAwaitPromise(true) // 等待Promise解析
			},
		),
	)
	if err != nil {
		return "", "", err // 返回错误
	}

	// 解析URL以获取域名
	domain := ""
	u, err := url.Parse(rawUrl)
	if err == nil {
		domain = u.Scheme + "://" + u.Host
	}

	// 设置HTML转Markdown的选项
	options := &md.Options{
		GetAbsoluteURL: func(selec *goquery.Selection, rawURL string, _ string) string {
			if !rewiseDomain {
				return rawUrl
			}
			if strings.HasPrefix(rawUrl, "/") {
				return domain + rawUrl
			}
			return rawUrl
		},
	}

	// 创建HTML到Markdown的转换器
	converter := md.NewConverter("", true, options)
	content, ok := result["content"].(string)
	if !ok {
		return "", "", fmt.Errorf("content not found") // 如果内容未找到，返回错误
	}

	// 将HTML内容转换为Markdown
	content, err = converter.ConvertString(content)
	if err != nil {
		return "", "", err // 转换失败，返回错误
	}

	// 获取标题
	title, ok := result["title"].(string)
	if !ok {
		title = "" // 如果标题未找到，设置为空字符串
	}

	return title, content, nil // 返回标题、内容和nil错误
}
