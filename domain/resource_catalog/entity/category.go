package entity

import "time"

// Category 类别实体
type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	ParentID    int64     `json:"parent_id"`
	Level       int       `json:"level"`
	Sort        int       `json:"sort"`
	Description string    `json:"description"`
	Status      int       `json:"status"` // 1:启用 0:禁用
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// IsActive 是否激活
func (c *Category) IsActive() bool {
	return c.Status == 1
}

// Validate 验证
func (c *Category) Validate() error {
	if c.Name == "" {
		return ErrCategoryNameEmpty
	}
	if c.Code == "" {
		return ErrCategoryCodeEmpty
	}
	return nil
}
