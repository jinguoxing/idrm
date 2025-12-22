package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/conf"

	"idrm/app/api/resource_catalog"
	"idrm/app/api/resource_catalog/handler"
	"idrm/domain/resource_catalog/service"
	"idrm/infrastructure/persistence/mysql"
	"idrm/pkg/config"
	"idrm/pkg/middleware"
)

var configFile = flag.String("f", "etc/api-server.yaml", "the config file")

func main() {
	flag.Parse()

	// 加载配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化数据库连接
	dbMap := make(map[string]*sql.DB)

	// 资源目录数据库
	catalogDB, err := initDB(c.DataSources.ResourceCatalog)
	if err != nil {
		log.Fatalf("Failed to init resource catalog database: %v", err)
	}
	dbMap["resource_catalog"] = catalogDB
	defer catalogDB.Close()

	// 数据视图数据库
	dataViewDB, err := initDB(c.DataSources.DataView)
	if err != nil {
		log.Fatalf("Failed to init data view database: %v", err)
	}
	dbMap["data_view"] = dataViewDB
	defer dataViewDB.Close()

	// 数据理解数据库
	dataUnderstandingDB, err := initDB(c.DataSources.DataUnderstanding)
	if err != nil {
		log.Fatalf("Failed to init data understanding database: %v", err)
	}
	dbMap["data_understanding"] = dataUnderstandingDB
	defer dataUnderstandingDB.Close()

	// 初始化仓储和服务
	categoryRepo := mysql.NewCategoryRepository(catalogDB)
	categoryService := service.NewCategoryService(categoryRepo)

	// 初始化处理器
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// 创建路由
	mux := http.NewServeMux()

	// 注册路由
	resource_catalog.RegisterRoutes(mux, categoryHandler)

	// 添加中间件
	var finalHandler http.Handler = mux
	finalHandler = middleware.CorsMiddleware(c.Cors.AllowOrigins, c.Cors.AllowMethods, c.Cors.AllowHeaders)(finalHandler)

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	log.Printf("Starting API server at %s", addr)

	if err := http.ListenAndServe(addr, finalHandler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// initDB 初始化数据库连接
func initDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, cfg.Source)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
