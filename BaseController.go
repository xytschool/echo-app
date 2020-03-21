package echoapp

import (
	echoapp_util "github.com/gw123/echo-app/util"
	"github.com/labstack/echo"
	"net/http"
)

const (
	Error_NotFound      = 400
	Error_NoAuth        = 401
	Error_DBError       = 402
	Error_CacheError    = 403
	Error_ArgumentError = 404
	Error_NotAllow      = 405
	Error_EtcdError     = 406
)

type Response struct {
	ErrorCode  int         `json:"code"`
	Msg        string      `json:"msg"`
	InnerError string      `json:"inner_error"`
	Data       interface{} `json:"data"`
}

type BaseController struct {
}

func (b *BaseController) Success(ctx echo.Context, data interface{}) error {
	response := Response{
		ErrorCode: 0,
		Msg:       "success",
		Data:      data,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (b *BaseController) Fail(ctx echo.Context, errcode int, msg string, innerErr error) error {
	innerErrStr := ""
	if innerErr != nil {
		innerErrStr = innerErr.Error()
		echoapp_util.ExtractEntry(ctx).WithError(innerErr).Info("requestFail")
	}

	response := Response{
		ErrorCode:  errcode,
		Msg:        msg,
		InnerError: innerErrStr,
	}
	return ctx.JSON(http.StatusOK, response)
}