package xhttp

import (
	"fmt"
	"net/http"

	"git.changjing.com.cn/zhongtai/yijing-common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

//http方法
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) int {
	if err == nil {
		// body, _ := json.Marshal(resp)
		// logx.WithContext(r.Context()).Info("resp:", string(body))
		//成功返回
		httpx.WriteJson(w, http.StatusOK, ResponsResp(errorx.SuccessCode, errorx.SuccessMsg, resp))
		return http.StatusOK
	} else {
		var (
			errCode int    = errorx.ErrCodeBadReq
			errMsg  string = err.Error()
		)
		switch err.(type) {
		case *errorx.CodeError:
			e := err.(*errorx.CodeError)
			errCode = e.Code()
			errMsg = e.Msg()
		default:
			fmt.Printf("err %+v", err)
			if gstatus, ok := status.FromError(errors.Cause(err)); ok {
				// grpc err错误
				errMsg = gstatus.Message()
			}
		}
		// logx.WithContext(r.Context()).Errorf("HttpResult err:%+v", err)
		httpx.WriteJson(w, http.StatusOK, ResponsResp(errCode, errMsg, resp))

		return errCode
	}
}

func HttpFileStream(r *http.Request, w http.ResponseWriter, err error, path string, name string) {
	if err != nil {
		var (
			errCode int    = errorx.ErrCodeApiServer
			errMsg  string = errorx.ErrMsgApiServer
		)
		//错误返回
		if e, ok := err.(*errorx.CodeError); ok {
			//自定义CodeError
			errCode = e.Code()
			errMsg = e.Msg()
		} else {
			originErr := errors.Cause(err) // err类型
			if gstatus, ok := status.FromError(originErr); ok {
				// grpc err错误
				errMsg = gstatus.Message()
			}
		}
		// logx.WithContext(r.Context()).Error("【GATEWAY-SRV-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusOK, ResponsResp(errCode, errMsg, nil))
		return
	}
	content := fmt.Sprintf("attachment;filename=%s", name)
	w.Header().Set("Content-Disposition", content)
	http.ServeFile(w, r, path)
}
