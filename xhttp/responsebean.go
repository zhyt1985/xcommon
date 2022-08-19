package xhttp

import "github.com/coolwxb/xcommon/xerr"

type ResponseSuccessBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{xerr.SuccessCode, xerr.SuccessMsg, data}
}

type ResponseErrorBean struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func Error(errCode int, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
