package logic

import (
	"context"

	"{{MODULE_PATH}}/api/internal/svc"
	"{{MODULE_PATH}}/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HealthCheckLogic 健康检查逻辑
type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewHealthCheckLogic 创建健康检查逻辑实例
func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HealthCheck 执行健康检查
func (l *HealthCheckLogic) HealthCheck() (resp *types.HealthResp, err error) {
	return &types.HealthResp{
		Status:  "ok",
		Version: "v1.0.0",
	}, nil
}
