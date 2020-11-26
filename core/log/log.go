package log

import (
	"context"
	"sfgo/core/log/base"
	"sfgo/core/log/zap"
)

var (
	// 确保不为空
	DefaultLogger base.ILogger = zap.New()
)

// 通用的初始化方法。
func InitLogger(logger base.ILogger) {
	DefaultLogger = logger
}

// 以下为log模块的输出方法。
func Debug(ctx context.Context, format string, args ...interface{}) {
	DefaultLogger.WithFields(Caller()).Debug(ctx, format, args...)
}

func Info(ctx context.Context, format string, args ...interface{}) {
	DefaultLogger.WithFields(Caller()).Info(ctx, format, args...)
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	DefaultLogger.WithFields(Caller()).Warn(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	DefaultLogger.WithFields(Caller()).Error(ctx, format, args...)
}

func Fatal(ctx context.Context, format string, args ...interface{}) {
	DefaultLogger.WithFields(Caller()).Fatal(ctx, format, args...)
}

func Panic(ctx context.Context, format string, args ...interface{}) {
	DefaultLogger.WithFields(Caller()).Panic(ctx, format, args...)
}

func WithFields(mapFields map[string]interface{}) base.ILogger {
	return DefaultLogger.WithFields(mapFields).WithFields(Caller())
}

func GetLevel() string {
	return DefaultLogger.GetLevel()
}

func GetLogger() base.ILogger {
	return DefaultLogger
}
