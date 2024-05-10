package middleware

import (
	"MyGo-scraper/common"
	"MyGo-scraper/common/logger"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		trace := logger.NewTraceLogger(c)
		trace.SetIp(c.RealIP())

		c.Set(common.CTX_TRACE_LOGGER, trace)
		c.Response().Header().Set("X-Trace-Id", trace.RequestId)

		if err := next(c); err != nil {
			c.Error(err)
		}

		trace.Write()
		return nil
	}
}
