package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// Producer Kafka生产者
type Producer struct {
	writer *kafka.Writer
}

// NewProducer 创建Kafka生产者
func NewProducer(brokers []string, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{
		writer: writer,
	}
}

// SendMessage 发送消息
func (p *Producer) SendMessage(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	return p.writer.WriteMessages(ctx, message)
}

// SendMessages 批量发送消息
func (p *Producer) SendMessages(ctx context.Context, messages []kafka.Message) error {
	return p.writer.WriteMessages(ctx, messages...)
}

// Close 关闭生产者
func (p *Producer) Close() error {
	return p.writer.Close()
}

// Consumer Kafka消费者
type Consumer struct {
	reader  *kafka.Reader
	handler MessageHandler
}

// MessageHandler 消息处理器接口
type MessageHandler interface {
	Handle(ctx context.Context, message kafka.Message) error
}

// NewConsumer 创建Kafka消费者
func NewConsumer(brokers []string, groupID, topic string, handler MessageHandler) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &Consumer{
		reader:  reader,
		handler: handler,
	}
}

// Start 启动消费者
func (c *Consumer) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			message, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Failed to read message: %v", err)
				continue
			}

			// 处理消息
			if err := c.handler.Handle(ctx, message); err != nil {
				log.Printf("Failed to handle message: %v", err)
				// 可以根据需要实现重试逻辑
			}
		}
	}
}

// Close 关闭消费者
func (c *Consumer) Close() error {
	return c.reader.Close()
}
