package main

import (
	"MyGo-scraper/common/config"
	"MyGo-scraper/common/initializer"
	"MyGo-scraper/common/logger"
)

func main() {
	logger.Init()
	config.Init()
	initializer.Init()

}
