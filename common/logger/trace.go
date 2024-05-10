package logger

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

type TraceLogger struct {
	ctx       echo.Context
	IP        string
	Uid       string
	RequestId string
	RequestTs int64
	Request   any
	Response  any
	BizCode   int
	BizMsg    string
	Params    []any
	Logger    *logrus.Entry
}

func NewTraceLogger(c echo.Context) *TraceLogger {
	// 生成随机ID
	requestId := xid.New().String()

	logger := Logger.WithFields(logrus.Fields{
		"reqid": requestId,
	})

	return &TraceLogger{
		ctx:       c,
		RequestId: requestId,
		RequestTs: time.Now().UnixMilli(),
		Params:    make([]any, 0),
		Logger:    logger,
	}
}

func (t *TraceLogger) SetIp(ip string) {
	t.IP = ip
}

func (t *TraceLogger) SetUid(uid string) {
	t.Uid = uid
}

func (t *TraceLogger) SetBizRequest(req any) {
	t.Request = req
}

func (t *TraceLogger) SetBizResponse(res any) {
	t.Response = res
}

func (t *TraceLogger) Trace(key string, value any) {
	t.Params = append(t.Params, key, value)
}

func (t *TraceLogger) Tracef(key string, format string, args ...any) {
	t.Params = append(t.Params, key, fmt.Sprintf(format, args...))
}

func (t *TraceLogger) Write() {
	duration := time.Now().UnixMilli() - t.RequestTs
	req := t.ctx.Request()

	// 构建基本的日志信息
	fields := logrus.Fields{
		"method":   req.Method,
		"path":     req.URL.Path,
		"uid":      t.Uid,
		"ip":       t.IP,
		"biz_code": t.BizCode,
		"biz_msg":  t.BizMsg,
		"duration": duration,
	}

	// 如果有业务请求数据，添加到日志字段中
	if t.Request != nil {
		fields["biz_req"] = t.Request
	}
	// 如果有业务响应数据，添加到日志字段中
	if t.Response != nil {
		fields["biz_res"] = t.Response
	}

	// ["key1", "value1", "key2", "value2", ...] 每两个进行处理
	// 将额外的参数添加到日志字段中
	for i := 0; i < len(t.Params); i += 2 {
		if i+1 < len(t.Params) {
			key, ok := t.Params[i].(string)
			if ok {
				fields[key] = t.Params[i+1]
			}
		}
	}

	// 使用logrus的WithFields方法记录日志
	t.Logger.WithFields(fields).Info("TraceLogger process")
}
