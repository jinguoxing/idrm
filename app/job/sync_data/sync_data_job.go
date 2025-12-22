package sync_data

import (
	"context"
	"log"
	"time"

	"idrm/domain/resource_catalog/service"
)

// SyncDataJob 数据同步任务
type SyncDataJob struct {
	categoryService *service.CategoryService
}

// NewSyncDataJob 创建数据同步任务
func NewSyncDataJob(categoryService *service.CategoryService) *SyncDataJob {
	return &SyncDataJob{
		categoryService: categoryService,
	}
}

// Run 执行任务
func (j *SyncDataJob) Run(ctx context.Context) error {
	log.Printf("[SyncDataJob] Starting sync data job at %s", time.Now().Format(time.RFC3339))

	// TODO: 实现数据同步逻辑
	// 1. 从外部系统获取数据
	// 2. 比对差异
	// 3. 更新本地数据
	// 4. 发送同步完成事件

	// 示例：获取所有类别
	categories, err := j.categoryService.GetTree(ctx)
	if err != nil {
		log.Printf("[SyncDataJob] Failed to get categories: %v", err)
		return err
	}

	log.Printf("[SyncDataJob] Synced %d categories", len(categories))

	log.Printf("[SyncDataJob] Completed sync data job at %s", time.Now().Format(time.RFC3339))
	return nil
}

// GetCron 获取cron表达式
func (j *SyncDataJob) GetCron() string {
	return "0 */5 * * * *" // 每5分钟执行一次
}

// GetName 获取任务名称
func (j *SyncDataJob) GetName() string {
	return "SyncDataJob"
}
