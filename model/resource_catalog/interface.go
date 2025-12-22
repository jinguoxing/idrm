package resource_catalog

import "context"

// CategoryModel 类别仓储接口（统一抽象）
// sqlx和gorm都需要实现此接口
type CategoryModel interface {
	Insert(ctx context.Context, data *Category) error
	FindOne(ctx context.Context, id int64) (*Category, error)
	FindByCode(ctx context.Context, code string) (*Category, error)
	Update(ctx context.Context, data *Category) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, page, pageSize int) ([]*Category, int64, error)
	FindByParentId(ctx context.Context, parentId int64) ([]*Category, error)
}
