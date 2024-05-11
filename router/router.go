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

	// CORS 跨域设置
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
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) (bool, error) {
			return true, nil
		},
		MaxAge: 7200, // 2 hour
	}))

	// Router 注册
	initRoutes(e)

	// 启动HTTP服务器
	e.Logger.Fatal(
		e.Start(":" + strconv.Itoa(config.AppConfig.HttpPort)),
	)
}

// 统一路由注册
var registry []func(e *echo.Echo)

func RegisterRoute(registerFunc func(e *echo.Echo)) {
	registry = append(registry, registerFunc)
}

func initRoutes(e *echo.Echo) {
	for _, register := range registry {
		register(e)
	}
}
