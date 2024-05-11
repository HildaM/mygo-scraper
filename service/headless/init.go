package headless

import (
	"MyGo-scraper/common/config"
	"MyGo-scraper/common/initializer"
	l "MyGo-scraper/common/logger"
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/chromedp/chromedp"
)

var (
	allocCtx      context.Context
	readabilityJS string
)

func initHeadless() error {
	cs := config.AppConfig.Chrome
	if cs.RemoteUrl != "" {
		l.Logger.Info(fmt.Printf("use remote chrome: %s", cs.RemoteUrl))

		ctx, _ := chromedp.NewRemoteAllocator(context.Background(), cs.RemoteUrl)
		allocCtx = ctx

	} else {
		// chromedp自定义配置
		options := []chromedp.ExecAllocatorOption{
			chromedp.Headless,
			chromedp.DisableGPU,
			chromedp.Flag("blink-settings", "imagesEnabled=true"),
		}
		if cs.ExecPath != "" {
			l.Logger.Info(fmt.Sprintf("use local chrome: %s", cs.ExecPath))
			options = append(options, chromedp.ExecPath(cs.ExecPath))
		} else {
			l.Logger.Info("use default chrome")
		}

		// chromedp默认配置
		options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

		// http proxy
		if config.AppConfig.HttpProxy != "" {
			if _, err := url.Parse(config.AppConfig.HttpProxy); err != nil {
				return err
			}

			l.Logger.Info(fmt.Sprintf("use http proxy: %s", config.AppConfig.HttpProxy))
			options = append(options, chromedp.ProxyServer(config.AppConfig.HttpProxy))
		}

		ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)
		allocCtx = ctx
	}

	js, err := os.ReadFile("assets/Readability.js")
	if err != nil {
		return fmt.Errorf("failed to read JavaScript file: %v", err)
	}

	readabilityJS = string(js)
	return nil
}

func init() {
	initializer.Register("headless", initHeadless)
}
