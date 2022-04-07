package xerr

var message map[int]string

//成功返回
const (
	SuccessCode = 200
	SuccessMsg  = "OK"
	OK  = 200
)

const (
	ErrCodeApiServer      = 10001 //api服务异常
	ErrCodeRpcServer      = 10002 //rpc服务异常
	ErrCodeDataBase       = 10003 //数据库异常
	ErrCodeBadReq         = 10004 //错误的请求
	ErrCodeToken          = 10005 //token错误
	ErrCodeTokenInvalid   = 10006 //token失效
	ErrCodeParams         = 10007 //参数错误
	ErrCodeParamsInvalid  = 10008 //参数非法
	ErrCodeRedis          = 10009 //连接异常
	ErrCodeDataNotFound   = 10010 //数据不存在
	ErrCodeDataRepeat     = 10011 //数据重复
	ErrCodeDataPermission = 10012 //无权限

	// 文件相关
	ErrCodeFileRead     = 10012 //文件读取失败
	ErrCodeFileUpload   = 10013 //文件上传失败
	ErrCodeFileNotExist = 10014 //文件不存在
	ErrCodeFileDown     = 10015 //文件下载失败
	ErrCodeFileSave     = 10016 //文件生成失败
	// 用户相关
	ErrCodeUserNotFound   = 10017 // 用户不存在
	ErrCodeUserPasswd     = 10018 // 用户密码错误
	ErrCodeUserNameRepeat = 10019 // 用户名重复

	// CRUD操作
	ErrCodeInsert = 10020 // 新增失败
	ErrCodeUpdate = 10021 // 修改失败
	ErrCodeDelete = 10022 // 删除失败
)
const (
	ErrMsgRpcServer     = "rpc服务错误"
	ErrMsgApiServer     = "api服务错误"
	ErrMsgDataBase      = "数据库异常"
	ErrMsgToken         = "token错误"
	ErrMsgTokeInvalid   = "token失效"
	ErrMsgDataRepeat    = "数据已存在"
	ErrMsgDataNotFound  = "数据不存在"
	ErrMsgPwdInvalid    = "密码无效"
	ErrMsgFileRead      = "文件读取失败"
	ErrMsgFileUpload    = "文件上传失败"
	ErrMsgFileNotExist  = "文件不存在"
	ErrMsgFileDown      = "文件下载失败"
	ErrMsgFileSave      = "文件生成失败"
	ErrMsgParams        = "参数错误"
	ErrMsgDatPermission = "无数据权限"
)


func init() {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[ErrCodeBadReq] = "服务器繁忙,请稍后再试"
	message[ErrCodeParams] = "参数错误"

	message[ErrCodeApiServer] = "api服务异常"
	message[ErrCodeRpcServer] = "rpc服务异常"
	message[ErrCodeDataBase] = "数据库异常"
	message[ErrCodeBadReq] = "错误的请求"
	message[ErrCodeToken] = "token错误"
	message[ErrCodeTokenInvalid] = "token失效"
	message[ErrCodeParamsInvalid] = "参数非法"
	message[ErrCodeParams] = "参数错误"
	message[ErrCodeRedis] = "Redis异常"
	message[ErrCodeDataNotFound] = "数据不存在"
	message[ErrCodeDataRepeat] = "数据重复"
	message[ErrCodeFileRead] = "文件读取失败"
	message[ErrCodeFileUpload] = "文件上传失败"
	message[ErrCodeFileNotExist] = "文件不存在"
	message[ErrCodeFileDown] = "文件下载失败"
	message[ErrCodeFileSave] = "文件生成失败"
	message[ErrCodeBadReq] = "服务器繁忙,请稍后再试"
	message[ErrCodeParams] = "参数错误"

	message[ErrCodeUserNotFound] = "用户不存在"
	message[ErrCodeUserPasswd] = "密码错误"
	message[ErrCodeUserNameRepeat] = "用户名重复"

	message[ErrCodeInsert] = "新增失败"
	message[ErrCodeUpdate] = "修改失败"
	message[ErrCodeDelete] = "删除失败"

}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器繁忙,请稍后再试"
	}
}
