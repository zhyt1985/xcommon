package xhttp

import "net/http"

func RespJson(writer http.ResponseWriter, code int, data string) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write([]byte(data))
}
