package resource_catalog

import (
	"database/sql"

	"idrm/model/resource_catalog/gorm"
	"idrm/model/resource_catalog/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
	gormdb "gorm.io/gorm"
)

// NewCategoryModel 创建Category模型（自动选择ORM）
// 优先使用gorm（更强大），如果gorm不可用则降级到sqlx
func NewCategoryModel(sqlConn *sql.DB, gormDB *gormdb.DB) CategoryModel {
	// 优先使用gorm
	if gormDB != nil {
		logx.Info("Using GORM for CategoryModel")
		return gorm.NewCategoryDao(gormDB)
	}

	// 降级使用sqlx
	if sqlConn != nil {
		logx.Info("Using SQLx for CategoryModel (fallback)")
		return sqlx.NewCategoryModel(sqlConn)
	}

	panic("no database connection available for CategoryModel")
}
