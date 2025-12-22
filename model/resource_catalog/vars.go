package resource_catalog

import "errors"

// 错误定义
var (
	ErrNotFound          = errors.New("category not found")
	ErrCodeAlreadyExists = errors.New("category code already exists")
	ErrInvalidStatus     = errors.New("invalid status")
)

// 状态常量
const (
	StatusDisabled = 0
	StatusEnabled  = 1
)
