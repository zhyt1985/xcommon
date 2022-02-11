package xerr

import "fmt"

type CodeError struct {
	errCode int
	errMsg  string
}

//属性
func (e *CodeError) GetErrCode() int {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func New(errCode int, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

// func NewErrMsg(errMsg string) *CodeError {
// 	return &CodeError{errCode: ErrCodeBadReq, errMsg: errMsg}
// }
