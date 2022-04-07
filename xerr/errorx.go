package xerr

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



// NewError 返回新定义的error接口
func NewError(code int, msg string) ErrorInfo {
	return &CodeError{code, msg}
}
func NewRpcError(code int, msg string) error {
	return status.Error(codes.Code(code), msg)
}



// ErrParams 参数错误
func ErrParams(msg string) ErrorInfo {
	return NewError(ErrCodeParams, fmt.Sprintf("\"%s\"", msg))
}

// ErrParamsFormat 参数格式错误
func ErrParamsFormat(msg string) ErrorInfo {
	return NewError(ErrCodeParams, fmt.Sprintf("参数格式错误:\"%s\"", msg))
}

func ErrApiServer(msg string) ErrorInfo {
	return NewError(ErrCodeApiServer, fmt.Sprintf("%s", msg))
}

/**************rpc通用错误****************/
// ErrRpcDataBase 数据库异常
func ErrRpcDataBase(msg string) error {
	return NewRpcError(ErrCodeDataBase, fmt.Sprintf("数据库异常:%s", msg))
}

// ErrRpcDataNotFount 数据不存在
func ErrRpcDataNotFount(msg string) error {
	return NewRpcError(ErrCodeDataNotFound, msg)
}

// ErrRpcService rpc服务异常
func ErrRpcService(msg string) ErrorInfo {
	return NewError(ErrCodeRpcServer, fmt.Sprintf(ErrMsgRpcServer+" %s", msg))
}
