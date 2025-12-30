# Job Service

定时任务服务目录，用于处理周期性的后台任务。

## 目录结构

```
job/
├── {job_name}/
│   ├── {job_name}.go        # 任务入口
│   ├── etc/
│   │   └── {job_name}.yaml  # 配置文件
│   └── internal/
│       ├── config/
│       ├── logic/
│       └── svc/
```

## 任务类型

- **周期任务**：定时执行的任务（如每日数据同步）
- **一次性任务**：只执行一次的任务
- **Cron 任务**：基于 Cron 表达式的任务

## 使用示例

```go
// 使用 go-zero 的 cron
import "github.com/zeromicro/go-zero/core/cron"

func main() {
    c := cron.NewCron()
    c.AddFunc("0 0 * * *", func() {
        // 每天 0 点执行
    })
    c.Start()
}
```

## 注意事项

- 定时任务应该是幂等的
- 长时间运行的任务需要支持优雅退出
- 注意任务重叠（overlap）的处理
