package xhttp

type ResponseSuccessBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func ResponsResp(code int, msg string, data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{code, msg, data}
}
