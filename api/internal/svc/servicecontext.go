package svc

import (
	"database/sql"
	"fmt"

	"idrm/api/internal/config"
	"idrm/model/resource_catalog"
	"idrm/pkg/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	// Model层（使用接口类型，支持自动ORM选择）
	CategoryModel resource_catalog.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 1. 初始化 sqlx 连接（作为备用）
	var sqlConn *sql.DB
	dsn := buildDSN(c.DB.ResourceCatalog)
	conn := sqlx.NewMysql(dsn)
	var err error
	sqlConn, err = conn.RawDB()
	if err != nil {
		logx.Errorf("SQLx RawDB获取失败: %v", err)
	}

	// 2. 初始化 gorm 连接（优先）
	var gormDB *gorm.DB
	gormDB, err = db.InitGorm(c.DB.ResourceCatalog)
	if err != nil {
		// gorm初始化失败，记录日志，将使用sqlx作为降级
		logx.Errorf("GORM初始化失败: %v, 将使用SQLx作为降级方案", err)
	}

	// 3. 使用工厂自动选择ORM（gorm优先，sqlx降级）
	categoryModel := resource_catalog.NewCategoryModel(sqlConn, gormDB)

	return &ServiceContext{
		Config:        c,
		CategoryModel: categoryModel,
	}
}

// buildDSN 构建 sqlx 的 DSN
func buildDSN(cfg db.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)
}
