package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(w http.ResponseWriter, data interface{}) {
	resp := &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	httpx.OkJson(w, resp)
}

// Error 错误响应
func Error(w http.ResponseWriter, code int, msg string) {
	resp := &Response{
		Code: code,
		Msg:  msg,
	}
	httpx.OkJson(w, resp)
}
