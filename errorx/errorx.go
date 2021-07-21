package errorx

import (
	"git.changjing.com.cn/zhongtai/yijing-common/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DescodeRpcErr 解码
func DescodeRpcErr(e error) ErrorInfo {
	var (
		code int = -1
		msg  string
	)
	if se, ok := status.FromError(e); ok {
		msg = se.Proto().Message
		code, _ = utils.GetInt(se.Proto().Code)
	}
	return NewError(code, msg)
}

// NewRpcError 初始化rpc错误
func NewRpcError(code int, msg string) error {
	return status.Error(codes.Code(code), msg)
}

type ErrorInfo interface {
	error
	Code() int
	Msg() string
}

type CodeError struct {
	ErrCode int    `json:"code"`
	ErrMsg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.ErrMsg
}

// Code 错误编码
func (e *CodeError) Code() int {
	return e.ErrCode
}

// 错误消息
func (e *CodeError) Msg() string {
	return e.ErrMsg
}

// NewError 返回新定义的error接口
func NewError(code int, msg string) ErrorInfo {
	return &CodeError{code, msg}
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code(),
		Msg:  e.Msg(),
	}
}
