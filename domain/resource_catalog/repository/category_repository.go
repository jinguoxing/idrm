package repository

import (
	"context"

	"idrm/domain/resource_catalog/entity"
)

// CategoryRepository 类别仓储接口
type CategoryRepository interface {
	// Create 创建类别
	Create(ctx context.Context, category *entity.Category) error

	// Update 更新类别
	Update(ctx context.Context, category *entity.Category) error

	// Delete 删除类别
	Delete(ctx context.Context, id int64) error

	// FindByID 根据ID查询
	FindByID(ctx context.Context, id int64) (*entity.Category, error)

	// FindByCode 根据编码查询
	FindByCode(ctx context.Context, code string) (*entity.Category, error)

	// FindByParentID 根据父ID查询子类别
	FindByParentID(ctx context.Context, parentID int64) ([]*entity.Category, error)

	// List 分页查询
	List(ctx context.Context, page, pageSize int) ([]*entity.Category, int64, error)

	// FindAll 查询所有激活的类别
	FindAll(ctx context.Context) ([]*entity.Category, error)
}
