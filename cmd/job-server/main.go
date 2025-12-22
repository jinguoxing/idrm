package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/conf"

	"idrm/app/job/sync_data"
	"idrm/domain/resource_catalog/service"
	"idrm/infrastructure/persistence/mysql"
	"idrm/pkg/config"
)

var configFile = flag.String("f", "etc/job-server.yaml", "the config file")

func main() {
	flag.Parse()

	// 加载配置
	var c config.JobConfig
	conf.MustLoad(*configFile, &c)

	// 初始化数据库
	catalogDB, err := initDB(c.DataSources.ResourceCatalog)
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}
	defer catalogDB.Close()

	// 初始化服务
	categoryRepo := mysql.NewCategoryRepository(catalogDB)
	categoryService := service.NewCategoryService(categoryRepo)

	// 创建定时任务调度器
	cronScheduler := cron.New(cron.WithSeconds())

	// 注册数据同步任务
	if c.Jobs.SyncData.Enabled {
		syncJob := sync_data.NewSyncDataJob(categoryService)
		_, err := cronScheduler.AddFunc(c.Jobs.SyncData.Cron, func() {
			ctx := context.Background()
			if err := syncJob.Run(ctx); err != nil {
				log.Printf("Failed to run sync job: %v", err)
			}
		})
		if err != nil {
			log.Fatalf("Failed to add sync job: %v", err)
		}
		log.Printf("Registered sync job with cron: %s", c.Jobs.SyncData.Cron)
	}

	// TODO: 注册其他任务
	// if c.Jobs.Statistics.Enabled { ... }
	// if c.Jobs.Cleanup.Enabled { ... }

	// 启动调度器
	cronScheduler.Start()
	log.Println("Job server started")

	// 等待退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down job server...")
	cronScheduler.Stop()
	log.Println("Job server stopped")
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
