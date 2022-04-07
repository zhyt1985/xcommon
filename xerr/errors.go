package xerr


type ErrorInfo interface {
	error
	Code() int
	Msg() string
}

type CodeError struct {
	errCode int
	errMsg  string
}
type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//属性
func (e *CodeError) GetErrCode() int {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return e.errMsg
}

// Code 错误编码
func (e *CodeError) Code() int {
	return e.errCode
}

// 错误消息
func (e *CodeError) Msg() string {
	return e.errMsg
}


func New(errCode int, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrByErrCode(errCode int) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewModeErr(mode string, errCode int, errMsg error) *CodeError {
	if mode == "pro" {
		return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
	}
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode) + ": " + errMsg.Error()}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: ErrCodeBadReq, errMsg: errMsg}
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.errCode,
		Msg:  e.errMsg,
	}
}
