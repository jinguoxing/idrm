package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/conf"

	"idrm/app/consumer/catalog_consumer"
	"idrm/domain/resource_catalog/service"
	"idrm/infrastructure/mq/kafka"
	"idrm/infrastructure/persistence/mysql"
	"idrm/pkg/config"
)

var configFile = flag.String("f", "etc/consumer-server.yaml", "the config file")

func main() {
	flag.Parse()

	// 加载配置
	var c config.ConsumerConfig
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

	// 创建消费者
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// 启动目录消费者
	if catalogCfg, ok := c.Kafka.Consumers["Catalog"]; ok {
		catalogHandler := catalog_consumer.NewCatalogConsumer(categoryService)

		for _, topic := range catalogCfg.Topics {
			consumer := kafka.NewConsumer(
				c.Kafka.Brokers,
				catalogCfg.Group,
				topic,
				catalogHandler,
			)

			wg.Add(1)
			go func(topic string) {
				defer wg.Done()
				log.Printf("Starting consumer for topic: %s", topic)
				if err := consumer.Start(ctx); err != nil && err != context.Canceled {
					log.Printf("Consumer error for topic %s: %v", topic, err)
				}
			}(topic)
		}
	}

	// TODO: 启动其他消费者
	// if dataViewCfg, ok := c.Kafka.Consumers["DataView"]; ok { ... }

	log.Println("Consumer server started")

	// 等待退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down consumer server...")
	cancel()
	wg.Wait()
	log.Println("Consumer server stopped")
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
