package xhttp

import (
	"fmt"
	"github.com/coolwxb/xcommon/utils"
	"github.com/coolwxb/xcommon/xerr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

const (
	CtxRespCode = "respCode"
)

// HttpResult 返回结果
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) int {
	var (
		code int
	)
	if err == nil {
		code = http.StatusOK
		//成功返回
		httpx.WriteJson(w, code, Success(resp))
	} else {
		var (
			errCode int    = xerr.ErrCodeBadReq
			errMsg  string = err.Error()
		)
		switch err.(type) {
		case *xerr.CodeError:
			e := err.(*xerr.CodeError)
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		default:
			if gstatus, ok := status.FromError(errors.Cause(err)); ok {
				// grpc err错误
				errMsg = gstatus.Message()
			}
		}
		logx.WithContext(r.Context()).Errorf("【GATEWAY-SRV-ERR】 : %+v ", err)
		code = errCode
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
	}
	w.Header().Set(CtxRespCode, utils.GetString(code))
	return code
}

func HttpFileStream(r *http.Request, w http.ResponseWriter, err error, path string, name string) {
	if err != nil {
		var (
			errCode int    = xerr.ErrCodeApiServer
			errMsg  string = xerr.ErrMsgApiServer
		)
		//错误返回
		if e, ok := err.(*xerr.CodeError); ok {
			//自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			originErr := errors.Cause(err) // err类型
			if gstatus, ok := status.FromError(originErr); ok {
				// grpc err错误
				errMsg = gstatus.Message()
			}
		}
		logx.WithContext(r.Context()).Error("【GATEWAY-SRV-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
		return
	}
	content := fmt.Sprintf("attachment;filename=%s", name)
	w.Header().Set("Content-Disposition", content)
	http.ServeFile(w, r, path)
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.ErrCodeParams), err.Error())
	httpx.WriteJson(w, http.StatusOK, xerr.ErrParams(errMsg))
}
