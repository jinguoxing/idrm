package category

import "context"

// Model 定义类别仓储接口（统一抽象）
// sqlx和gorm都需要实现此接口
type Model interface {
	// 基础CRUD操作
	Insert(ctx context.Context, data *Category) (*Category, error)
	FindOne(ctx context.Context, id int64) (*Category, error)
	FindByCode(ctx context.Context, code string) (*Category, error)
	Update(ctx context.Context, data *Category) error
	Delete(ctx context.Context, id int64) error

	// 列表查询
	FindAll(ctx context.Context) ([]*Category, error)
	FindByParentId(ctx context.Context, parentId int64) ([]*Category, error)
	List(ctx context.Context, page, pageSize int) ([]*Category, int64, error)

	// 事务支持
	WithTx(tx interface{}) Model
	Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}
