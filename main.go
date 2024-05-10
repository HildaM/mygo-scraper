package main

import (
	"MyGo-scraper/common/config"
	"MyGo-scraper/common/initializer"
	"MyGo-scraper/common/logger"
	"MyGo-scraper/router"
)

func main() {
	logger.Init()
	config.Init()
	initializer.Init()
	router.Init()
}
