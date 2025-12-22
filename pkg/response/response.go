package response

import (
	"net/http"

	"idrm/pkg/errorx"
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
	WriteJSON(w, http.StatusOK, resp)
}

// Error 错误响应
func Error(w http.ResponseWriter, err error) {
	var code int
	var msg string

	if e, ok := err.(*errorx.CodeError); ok {
		code = e.GetCode()
		msg = e.GetMsg()
	} else {
		code = errorx.ErrCodeSystem
		msg = err.Error()
	}

	resp := &Response{
		Code: code,
		Msg:  msg,
	}

	WriteJSON(w, http.StatusOK, resp)
}

// SuccessWithMsg 带消息的成功响应
func SuccessWithMsg(w http.ResponseWriter, msg string, data interface{}) {
	resp := &Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	WriteJSON(w, http.StatusOK, resp)
}

// ErrorWithMsg 自定义错误消息响应
func ErrorWithMsg(w http.ResponseWriter, code int, msg string) {
	resp := &Response{
		Code: code,
		Msg:  msg,
	}
	WriteJSON(w, http.StatusOK, resp)
}

// WriteJSON 写JSON响应
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	// 这里需要使用json.Marshal序列化data并写入w
	// 为简化示例，实际使用时需要导入encoding/json
}
