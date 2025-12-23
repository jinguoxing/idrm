package telemetry

import (
	"context"

	"idrm/pkg/telemetry/audit"
	"idrm/pkg/telemetry/log"
	"idrm/pkg/telemetry/trace"

	"github.com/zeromicro/go-zero/core/logx"
)

// Init 初始化 Telemetry 系统（一站式初始化）
func Init(config Config) error {
	// 1. 初始化日志系统
	log.Init(config.Log, config.ServiceName)
	logx.Infof("Telemetry 初始化: %s v%s (%s)",
		config.ServiceName, config.ServiceVersion, config.Environment)

	// 2. 初始化链路追踪
	if err := trace.Init(config.Trace, config.ServiceName, config.ServiceVersion, config.Environment); err != nil {
		logx.Errorf("链路追踪初始化失败: %v", err)
		return err
	}

	// 3. 初始化审计日志
	audit.Init(config.Audit, config.ServiceName)

	logx.Info("Telemetry 系统初始化完成")
	return nil
}

// Close 关闭 Telemetry 系统
func Close(ctx context.Context) {
	logx.Info("正在关闭 Telemetry 系统...")

	// 关闭审计日志
	audit.Close()

	// 关闭链路追kick
	trace.Close(ctx)

	// 关闭日志系统（最后关闭）
	log.Close()
}
