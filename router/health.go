package router

import "github.com/labstack/echo/v4"

func init() {
	RegisterRoute(func(e *echo.Echo) {
		e.GET("/health", func(c echo.Context) error {
			c.String(200, "ok")
			return nil
		})
	})
}
