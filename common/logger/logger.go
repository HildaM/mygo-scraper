package logger

import (
	"MyGo-scraper/common"
	"MyGo-scraper/common/config"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Init() {
	rootDir := config.GetRootDir()

	// 创建一个新的logrus Logger实例
	Logger = logrus.New()

	// 设置日志输出格式为JSON
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 构造日志文件路径
	logFile := time.Now().Format("20060102") + ".log"
	logPath := filepath.Join(rootDir, "log", logFile)

	// 尝试打开或创建日志文件
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// 如果打开或创建文件失败，使用panic中断程序
		panic(err)
	}

	// 设置日志的输出为标准输出和文件
	Logger.SetOutput(io.MultiWriter(os.Stdout, file))

	// 设置日志级别为Info以上（默认级别）
	Logger.SetLevel(logrus.InfoLevel)
}

func GetTraceLogger(c echo.Context) *TraceLogger {
	trace, ok := c.Get(common.CTX_TRACE_LOGGER).(TraceLogger)
	if ok {
		return &trace
	}
	return nil
}
