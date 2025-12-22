package svc

import (
	"idrm/api/internal/config"
	"idrm/model/resource_catalog"
	gormResourceCatalog "idrm/model/resource_catalog/gorm"
	sqlxResourceCatalog "idrm/model/resource_catalog/sqlx"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	// Model层（使用接口类型，支持ORM切换）
	CategoryModel resource_catalog.CategoryModel
	// 其他model...
}

func NewServiceContext(c config.Config) *ServiceContext {
	var categoryModel resource_catalog.CategoryModel

	// 根据配置选择ORM
	if c.ORM == "gorm" {
		// 使用gorm
		db, err := gorm.Open(mysql.Open(c.DataSources.ResourceCatalog.Source), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		categoryModel = gormResourceCatalog.NewCategoryModel(db)
	} else {
		// 默认使用sqlx
		conn := sqlx.NewMysql(c.DataSources.ResourceCatalog.Source)
		categoryModel = sqlxResourceCatalog.NewCategoryModel(conn)
	}

	return &ServiceContext{
		Config:        c,
		CategoryModel: categoryModel,
	}
}
