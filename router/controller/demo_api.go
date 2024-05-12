package controller

import (
	"MyGo-scraper/common/logger"
	"MyGo-scraper/service/scrape"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetScrapeHTML(c echo.Context) error {
	urlParam := c.QueryParam("u")
	headless := c.QueryParam("headless") == "true"
	readability := c.QueryParam("readability") == "true"

	if urlParam == "" {
		c.Redirect(http.StatusFound, "https://github.com/HildaM/mygo-scraper")
		return nil
	}

	trace := logger.GetTraceLogger(c)
	if trace != nil {
		logger.Logger.Debug(fmt.Sprintf("into the trace: %v", trace))
		trace.Trace("url", urlParam)
	}

	// 创建爬虫
	s := scrape.NewScrape(headless, readability, true)
	res, err := s.Run(c, urlParam)

	if err != nil {
		c.Render(http.StatusOK, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title":   res.Title,
		"content": res.Content,
	})
}
