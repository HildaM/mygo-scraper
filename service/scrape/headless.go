package scrape

import (
	hl "MyGo-scraper/service/headless"
)

// headlessScrape 无头浏览器爬取网页信息
func (s *Scrape) headlessScrape(rawUrl string) (*ScrapeResult, error) {
	title, content, err := hl.Headless(rawUrl, s.readability, s.rewiseDomain)
	if err != nil {
		return nil, err
	}

	return &ScrapeResult{
		Url:     rawUrl,
		Title:   title,
		Content: content,
	}, nil
}
