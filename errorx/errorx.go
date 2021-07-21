package errorx

import (
	"git.changjing.com.cn/zhongtai/yijing-common/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DescodeRpcErr 解码
func DescodeRpcErr(e error) error {
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

type codeError struct {
	ErrCode int    `json:"code"`
	ErrMsg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *codeError) Error() string {
	return e.ErrMsg
}

// Code 错误编码
func (e *codeError) Code() int {
	return e.ErrCode
}

// 错误消息
func (e *codeError) Msg() int {
	return e.Msg()
}

// NewError 返回原有error接口
// 使用go-zero
func NewError(code int, msg string) error {
	return &codeError{code, msg}
}

// NewErrorInfo 返回新定义的error接口
// 使用普通项目框架
func NewErrorInfo(code int, msg string) ErrorInfo {
	return &codeError{code, msg}
}
