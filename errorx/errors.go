package errorx

import (
	"fmt"
	"runtime"
	// "git.changjing.com.cn/zhongtai/yijing-common/errorx"
)

//成功返回
const (
	SuccessCode = 200
	SuccessMsg  = "OK"
)

const (
	ErrCodeApiServer     = 10001 //api/v1服务异常
	ErrCodeRpcServer     = 10002 //rpc服务异常
	ErrCodeDataBase      = 10003 //数据库异常
	ErrCodeBadReq        = 10004 //错误的请求
	ErrCodeToken         = 10005 //token错误
	ErrCodeTokenInvalid  = 10006 //token失效
	ErrCodeParams        = 10007 //参数错误
	ErrCodeParamsInvalid = 10008 //参数非法
	ErrCodeRedis         = 10009 //连接异常
	ErrCodeDataNotFound  = 10010 //数据不存在
	ErrCodeDataRepeat    = 10011 //数据重复
	ErrCodeFileRead      = 10012 //文件读取失败
	ErrCodeFileUpload    = 10013 //文件上传失败
	ErrCodeFileNotExist  = 10014 //文件不存在
	ErrCodeFileDown      = 10015 //文件下载失败
	ErrCodeFileSave      = 10016 //文件生成失败
	ErrCodePhoneNotFound = 10017 //手机号不存在
	ErrCodePasswordErr   = 10018 //密码不正
	ErrCodeParamsFormat  = 10019 //参数格式错误
	//审批流10901开始
	ErrWorkFlowFail = 10100 //工作流信息
)
const (
	ErrMsgRpcServer    = "rpc服务错误"
	ErrMsgApiServer    = "api服务错误"
	ErrMsgDataBase     = "数据库异常"
	ErrMsgToken        = "token错误"
	ErrMsgTokeInvalid  = "token失效"
	ErrMsgDataRepeat   = "数据已存在"
	ErrMsgDataNotFound = "数据不存在"
	ErrMsgPwdInvalid   = "密码无效"
	ErrMsgFileRead     = "文件读取失败"
	ErrMsgFileUpload   = "文件上传失败"
	ErrMsgFileNotExist = "文件不存在"
	ErrMsgFileDown     = "文件下载失败"
	ErrMsgFileSave     = "文件生成失败"
	ErrMsgParams       = "参数错误"
	/**************用户*****************/
	ErrMsgUserNotFount = "用户不存在"
	/**************客户****************/
	ErrMsgCustomerUploadFail = "客户信息上传失败"

	ErrMsgProjectBusinessUploadFail = "商企项目上传失败"
	ErrMsgProjectPublicUploadFail   = "公建项目上传失败"

	/***************审批流*****************/
	ErrMsgWfProcessStart = "开启审批流失败"
	ErrMsgWfStore        = "工作流入库失败"
)

//用户模块 2开头
const (
	ErrCodeUserNotFound = 20001
)

/*-----------------------------错误方法--------------------------------------------------- */

// ErrPhoneNotFound ...
func ErrUserNotFound() ErrorInfo {
	return NewError(ErrCodeToken, ErrMsgUserNotFount)
}

// ErrPhoneNotFound ...
func ErrPhoneNotFound(msg string) ErrorInfo {
	return NewError(ErrCodePhoneNotFound, msg)
}

// ErrPasswordErr ...
func ErrPasswordErr(msg string) ErrorInfo {
	return NewError(ErrCodePasswordErr, msg)
}

// ErrParamsConvert 参数错误
func ErrParams(msg string) ErrorInfo {
	return NewError(ErrCodeParams, msg)
}

// ErrParamsFormat 参数格式错误
func ErrParamsFormat(msg string) ErrorInfo {
	return NewError(ErrCodeParamsFormat, fmt.Sprintf("参数格式错误:\"%s\"", msg))
}

// ErrParamsValueInvalid 参数无效
func ErrParamsValueInvalid(name string, value interface{}) ErrorInfo {
	return NewError(ErrCodeParamsInvalid, fmt.Sprintf("非法的参数值，参数 \"%s\"，值 \"%v\"", name, value))
}

// ErrRpcService rpc服务异常
func ErrRpcService(msg string) ErrorInfo {
	return NewError(ErrCodeRpcServer, fmt.Sprintf(ErrMsgRpcServer+" %s", msg))
}

// ErrApiService api服务异常
func ErrApiService(msg string) ErrorInfo {
	return NewError(ErrCodeApiServer, msg)
}

// ErrDataNotFound 数据不存在
func ErrDataNotFound() ErrorInfo {
	return NewError(ErrCodeDataNotFound, ErrMsgDataNotFound)
}

// ErrDataBase 数据库异常
func ErrDataBase() ErrorInfo {
	return NewError(ErrCodeDataBase, ErrMsgDataBase)
}

// ErrDatabase 数据库异常
func ErrDatabase(msg string) ErrorInfo {
	return NewError(ErrCodeDataBase, fmt.Sprintf("%s：%s", ErrMsgDataBase, msg))
}

// ErrDataRepeat 数据已存在
func ErrDataRepeat() ErrorInfo {
	return NewError(ErrCodeDataRepeat, ErrMsgDataRepeat)
}

// ErrToken token错误
func ErrToken() ErrorInfo {
	return NewError(ErrCodeToken, ErrMsgToken)
}

// ErrTokenInvalid token失效
func ErrTokenInvalid() ErrorInfo {
	return NewError(ErrCodeTokenInvalid, ErrMsgTokeInvalid)
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

// ErrRpcDataRepeat 数据已存在
func ErrRpcDataRepeat() error {
	return NewRpcError(ErrCodeDataRepeat, ErrMsgDataRepeat)
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

// NewApiError new
func NewApiError(code int, msg string) ErrorInfo {
	return NewError(code, msg)
}

// NewApiError new
func NewRpcErr(code int, msg string) error {
	return NewRpcError(code, msg)
}

// DescodeRpcError 解码err
func DescodeRpcError(err error) ErrorInfo {
	return DescodeRpcErr(err)
}
