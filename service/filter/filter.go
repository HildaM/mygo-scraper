package filter

import (
	"MyGo-scraper/common/config"
	"net/url"
	"strings"
)

type Filter struct {
	hostMap     map[string]bool
	blockString []string
}

func NewFilter(conf config.FilterRule) *Filter {
	hostMap := make(map[string]bool)
	for _, h := range conf.Host {
		hostMap[h] = true
	}

	return &Filter{
		hostMap:     hostMap,
		blockString: conf.BlockString,
	}
}

// Pass 过滤检查
func (f *Filter) Pass(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if _, ok := f.hostMap[u.Host]; ok {
		return false
	}

	return true
}

// ContentFilter 移除违规词
func (f *Filter) ContentFilter(content string) string {
	for _, b := range f.blockString {
		if b != "" {
			content = strings.Replace(content, b, "", -1)
		}
	}

	return content
}
