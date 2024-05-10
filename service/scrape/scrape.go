package scrape

import "github.com/labstack/echo/v4"

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

func (s *Scrape) Run(ctx echo.Context, rawUrl string) (*ScrapeResult, error) {
	return nil, nil
}
