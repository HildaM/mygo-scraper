package router

import (
	"MyGo-scraper/common/config"
	"MyGo-scraper/common/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	// 配置logger
	e.Logger = logger.GetEchoLogger()
	e.Use(logger.Hook())

	// 使用Recovery中间件来恢复panic
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentLength,
			echo.HeaderContentType,
		},
	}))

	// TODO Router 设置

	// 启动HTTP服务器
	e.Logger.Fatal(
		e.Start(":" + strconv.Itoa(config.AppConfig.HttpPort)),
	)
}
