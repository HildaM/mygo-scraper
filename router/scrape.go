package router

import (
	"MyGo-scraper/router/controller"
	"MyGo-scraper/router/middleware"

	"github.com/labstack/echo/v4"
)

func init() {
	RegisterRoute(func(e *echo.Echo) {
		scraperRouter(e)
	})
}

func scraperRouter(e *echo.Echo) {
	v1 := e.Group("/v1")
	// 对v1下的请求做统一的日志收集
	v1.Use(middleware.LoggerMiddleware)

	v1.POST("/scrape", controller.PostScrape)
}
