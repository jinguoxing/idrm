package entity

import "errors"

// 实体层错误定义
var (
	ErrCategoryNameEmpty  = errors.New("类别名称不能为空")
	ErrCategoryCodeEmpty  = errors.New("类别编码不能为空")
	ErrDirectoryNameEmpty = errors.New("目录名称不能为空")
)
