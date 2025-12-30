# Consumer Service

消息消费者服务目录，用于处理消息队列中的消息。

## 目录结构

```
consumer/
├── {consumer_name}/
│   ├── {consumer_name}.go        # 消费者入口
│   ├── etc/
│   │   └── {consumer_name}.yaml  # 配置文件
│   └── internal/
│       ├── config/
│       ├── logic/
│       └── svc/
```

## 支持的消息队列

- **Kafka**：高吞吐量消息队列
- **RabbitMQ**：AMQP 协议消息队列
- **Redis**：轻量级消息队列（Stream/Pub-Sub）

## 使用示例

```go
// Kafka Consumer
import "github.com/zeromicro/go-queue/kq"

func main() {
    consumer := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(func(k, v string) error {
        // 处理消息
        return nil
    }))
    defer consumer.Stop()
    consumer.Start()
}
```

## 注意事项

- 消息处理应该是幂等的
- 注意消费者组的配置
- 合理设置重试策略
- 考虑死信队列（DLQ）
