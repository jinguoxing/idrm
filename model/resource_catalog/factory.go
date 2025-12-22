package resource_catalog

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	gormdb "gorm.io/gorm"
)

// CategoryModelFactory 类别模型工厂函数类型
type CategoryModelFactory func(interface{}) CategoryModel

var (
	gormCategoryFactory CategoryModelFactory
	sqlxCategoryFactory CategoryModelFactory
)

// RegisterGormCategoryFactory 注册gorm工厂（由gorm包调用）
func RegisterGormCategoryFactory(factory CategoryModelFactory) {
	gormCategoryFactory = factory
}

// RegisterSqlxCategoryFactory 注册sqlx工厂（由sqlx包调用）
func RegisterSqlxCategoryFactory(factory CategoryModelFactory) {
	sqlxCategoryFactory = factory
}

// NewCategoryModel 创建Category模型（自动选择ORM）
// 优先使用gorm（更强大），如果gorm不可用则降级到sqlx
func NewCategoryModel(sqlConn *sql.DB, gormDB *gormdb.DB) CategoryModel {
	// 优先使用gorm
	if gormDB != nil && gormCategoryFactory != nil {
		logx.Info("Using GORM for CategoryModel")
		return gormCategoryFactory(gormDB)
	}

	// 降级使用sqlx
	if sqlConn != nil && sqlxCategoryFactory != nil {
		logx.Info("Using SQLx for CategoryModel (fallback)")
		return sqlxCategoryFactory(sqlConn)
	}

	panic("no database connection available for CategoryModel")
}
