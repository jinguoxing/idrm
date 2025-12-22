package errorx

// CodeError自定义错误
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() interface{} {
	return nil
}

// NewCodeError创建错误
func NewCodeError(code int, msg string) error {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}
