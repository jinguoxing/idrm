package catalog_consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"

	"idrm/domain/resource_catalog/service"
)

// CatalogEvent 目录事件
type CatalogEvent struct {
	EventType string      `json:"event_type"` // created, updated, deleted
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

// CatalogConsumer 目录消费者
type CatalogConsumer struct {
	categoryService *service.CategoryService
}

// NewCatalogConsumer 创建目录消费者
func NewCatalogConsumer(categoryService *service.CategoryService) *CatalogConsumer {
	return &CatalogConsumer{
		categoryService: categoryService,
	}
}

// Handle 处理消息
func (c *CatalogConsumer) Handle(ctx context.Context, message kafka.Message) error {
	log.Printf("Received message: key=%s, value=%s, partition=%d, offset=%d",
		string(message.Key),
		string(message.Value),
		message.Partition,
		message.Offset,
	)

	// 解析事件
	var event CatalogEvent
	if err := json.Unmarshal(message.Value, &event); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return err
	}

	// 根据事件类型处理
	switch event.EventType {
	case "created":
		return c.handleCreated(ctx, event)
	case "updated":
		return c.handleUpdated(ctx, event)
	case "deleted":
		return c.handleDeleted(ctx, event)
	default:
		log.Printf("Unknown event type: %s", event.EventType)
		return nil
	}
}

// handleCreated 处理创建事件
func (c *CatalogConsumer) handleCreated(ctx context.Context, event CatalogEvent) error {
	log.Printf("Handling created event: %+v", event)
	// TODO: 实现业务逻辑，比如同步到其他系统、发送通知等
	return nil
}

// handleUpdated 处理更新事件
func (c *CatalogConsumer) handleUpdated(ctx context.Context, event CatalogEvent) error {
	log.Printf("Handling updated event: %+v", event)
	// TODO: 实现业务逻辑
	return nil
}

// handleDeleted 处理删除事件
func (c *CatalogConsumer) handleDeleted(ctx context.Context, event CatalogEvent) error {
	log.Printf("Handling deleted event: %+v", event)
	// TODO: 实现业务逻辑
	return nil
}
