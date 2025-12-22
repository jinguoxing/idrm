# Consumer 层

消息队列消费服务目录。

## 用途

用于存放 Kafka、RabbitMQ 等消息队列的消费者代码。

## 结构示例

```
consumer/
├── catalog_consumer/     # 目录消费者
├── data_view_consumer/   # 数据视图消费者
└── ...
```

## 注意事项

- 按业务模块组织消费者
- 每个消费者独立目录
- 可复用 `model/` 和 `pkg/` 中的代码
