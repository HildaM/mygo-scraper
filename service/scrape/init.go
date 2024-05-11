package scrape

import (
	"MyGo-scraper/common/config"
	"MyGo-scraper/common/initializer"
	"net/http"
	"net/url"
	"time"
)

var (
	scrapeClient *http.Client
	filter       *Filter
)

func initScrape() error {
	scrapeClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:          30,                 // 设置最大空闲连接数为30
			IdleConnTimeout:       3600 * time.Second, // 设置空闲连接的超时时间为3600秒（1小时）
			ResponseHeaderTimeout: 360 * time.Second,  // 设置等待响应头的超时时间为360秒（6分钟）
			ExpectContinueTimeout: 360 * time.Second,  // 设置Expect: 100-continue的等待超时时间为360秒（6分钟）
			DisableCompression:    false,              // 设置是否禁用压缩，false表示不禁用，即允许压缩
		},
	}

	// 设置 http proxy
	if config.AppConfig.HttpProxy != "" {
		proxyUrl, err := url.Parse(config.AppConfig.HttpProxy)
		if err != nil {
			return err
		}
		scrapeClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	filter = NewFilter(config.AppConfig.Filter)
	return nil
}

func init() {
	initializer.Register("scrape", initScrape)
}
