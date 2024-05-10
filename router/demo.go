package router

import (
	"MyGo-scraper/common/logger"
	"MyGo-scraper/router/controller"

	"github.com/labstack/echo/v4"
)

func init() {
	RegisterRoute(func(e *echo.Echo) {
		// 在Echo中设置静态文件的目录
		e.Static("/", "static")

		e.GET("/", controller.GetScrapeHTML, logger.Hook())
	})
}
