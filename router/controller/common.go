package controller

import (
	"MyGo-scraper/common/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 响应码
const (
	CodeSuccess     int = 0
	CodeBadRequest  int = 400
	CodeServerError int = 500
)

// 状态信息
const (
	MsgSuccess = "ok"
	MsgFailed  = "failed"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func returnResponse(c echo.Context, code int, msg string, data any) {
	if trace := logger.GetTraceLogger(c); trace != nil {
		trace.BizCode = code
		trace.BizMsg = msg
	}

	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ReturnSuccess(c echo.Context, data any) {
	returnResponse(c, CodeSuccess, MsgSuccess, data)
}

func ReturnServerError(c echo.Context, err error) {
	returnResponse(c, CodeServerError, err.Error(), nil)
}

func ReturnBadRequest(c echo.Context, err error) {
	returnResponse(c, CodeBadRequest, err.Error(), nil)
}
