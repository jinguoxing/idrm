package category

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// Factory 类别模型工厂函数类型
type Factory func(interface{}) Model

var (
	gormFactory Factory
	sqlxFactory Factory
)

// RegisterGormFactory 注册gorm工厂（由gorm_dao.go调用）
func RegisterGormFactory(factory Factory) {
	gormFactory = factory
}

// RegisterSqlxFactory 注册sqlx工厂（由sqlx_model.go调用）
func RegisterSqlxFactory(factory Factory) {
	sqlxFactory = factory
}

// NewModel 创建Category模型（自动选择ORM）
// 优先使用gorm（更强大），如果gorm不可用则降级到sqlx
func NewModel(sqlConn *sql.DB, gormDB *gorm.DB) Model {
	// 优先使用gorm
	if gormDB != nil && gormFactory != nil {
		logx.Info("Using GORM for CategoryModel")
		return gormFactory(gormDB)
	}

	// 降级使用sqlx
	if sqlConn != nil && sqlxFactory != nil {
		logx.Info("Using SQLx for CategoryModel (fallback)")
		return sqlxFactory(sqlConn)
	}

	panic("no database connection available for CategoryModel")
}
