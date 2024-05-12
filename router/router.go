package router

import (
	"MyGo-scraper/common/config"
	"MyGo-scraper/common/logger"
	"io"
	"net/http"
	"strconv"
	"text/template"

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

	// 注册 HTML 渲染器
	e.Renderer = &HTMLTemplateRenderer{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}

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

// Html 渲染
// TemplateRenderer is a custom html/template renderer for Echo framework
type HTMLTemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *HTMLTemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
