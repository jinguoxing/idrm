package errorx

import "fmt"

// CodeError 业务错误
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// NewCodeError 创建业务错误
func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

// Error 实现 error 接口
func (e *CodeError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

// 常用错误码
const (
	// 系统错误 10000-19999
	CodeServerError = 10001

	// 业务错误 20000-29999
	CodeBizError = 20001

	// 验证错误 30000-39999
	CodeValidationError = 30001

	// 权限错误 40000-49999
	CodeUnauthorized = 40001
	CodeForbidden    = 40003
	CodeNotFound     = 40004
)

// 常用错误
var (
	ErrServerError     = NewCodeError(CodeServerError, "服务器内部错误")
	ErrUnauthorized    = NewCodeError(CodeUnauthorized, "未授权")
	ErrForbidden       = NewCodeError(CodeForbidden, "禁止访问")
	ErrNotFound        = NewCodeError(CodeNotFound, "资源不存在")
	ErrValidationError = NewCodeError(CodeValidationError, "参数验证失败")
)
