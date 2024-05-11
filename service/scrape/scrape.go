package scrape

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
)

type Scrape struct {
	rewiseDomain bool
	headless     bool
	readability  bool
	pipeline     []func(string) string
}

type ScrapeResult struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Error   string `json:"error,omitempty"`
}

func NewScrape(headless, readability, rewiseDomain bool) *Scrape {
	return &Scrape{
		headless:     headless,
		readability:  readability,
		rewiseDomain: rewiseDomain,
	}
}

func (s *Scrape) AddPipeline(fn func(string) string) {
	s.pipeline = append(s.pipeline, fn)
}

func (s *Scrape) Run(ctx echo.Context, rawUrl string) (*ScrapeResult, error) {
	// 1. 链接过滤
	if !filter.Pass(rawUrl) {
		return nil, fmt.Errorf("url not allowed: %s", rawUrl)
	}

	// 2. 爬取数据
	var res *ScrapeResult
	var err error

	if s.headless {
		res, err = s.headlessScrape(rawUrl)
	} else {
		res, err = s.directScrape(ctx, rawUrl)
	}
	if err != nil {
		return nil, err
	}

	// 3. 数据处理
	res.Content = filter.ContentFilter(res.Content)
	for _, fn := range s.pipeline {
		res.Content = fn(res.Content)
	}
	return res, nil
}

func (s *Scrape) BatchRun(ctx echo.Context, urlList []string) ([]ScrapeResult, error) {
	if len(urlList) == 0 {
		return nil, fmt.Errorf("url list is empty")
	}

	res := make([]ScrapeResult, len(urlList))

	var wg sync.WaitGroup
	for idx, url := range urlList {
		wg.Add(1)

		go func(idx int, url string) {
			defer wg.Done()

			res[idx] = ScrapeResult{Url: url} // 在并发的时候，使用拷贝赋值较为稳妥
			// 如果某个url爬取出现错误，立即recover，避免造成程序崩溃。同时记录error
			defer func() {
				if r := recover(); r != nil {
					res[idx].Error = fmt.Sprintf("panic: %v", r)
				}
			}()

			ret, err := s.Run(ctx, url)
			if err != nil {
				res[idx].Error = err.Error()
			}
			res[idx].Title = ret.Title
			res[idx].Content = ret.Content
		}(idx, url)
	}
	wg.Wait()

	return res, nil
}
