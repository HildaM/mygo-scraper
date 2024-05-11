package controller

import (
	"MyGo-scraper/service/scrape"
	"errors"

	"github.com/labstack/echo/v4"
)

type PostScrapeRequest struct {
	UrlList     []string `json:"url_list"`
	Headless    *bool    `json:"headless,omitempty"`
	Readability *bool    `json:"readability,omitempty"`
}

func PostScrape(c echo.Context) error {
	// 1. 解析请求参数
	req := &PostScrapeRequest{}
	if err := c.Bind(req); err != nil {
		ReturnBadRequest(c, err)
		return err
	}
	if len(req.UrlList) == 0 {
		err := errors.New("url list is empty")
		ReturnBadRequest(c, err)
		return err
	}

	// 2. 配置爬虫参数
	headless := true
	if req.Headless != nil {
		headless = *req.Headless
	}

	readability := true
	if req.Readability != nil {
		readability = *req.Readability
	}

	// 3. 执行爬虫逻辑
	s := scrape.NewScrape(headless, readability, true)
	res, err := s.BatchRun(c, req.UrlList)
	if err != nil {
		ReturnServerError(c, err)
		return err
	}

	ReturnSuccess(c, res)
	return nil
}
