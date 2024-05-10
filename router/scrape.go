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
	v1.Use(middleware.LoggerMiddleware)

	v1.POST("/scrape", controller.PostScrape)
}
