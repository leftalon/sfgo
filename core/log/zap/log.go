package zap

import (
	"context"
	"fmt"
	"sfgo/core/log/base"

	"go.uber.org/zap"
)

type zapLogger struct{}

func NewLogger() *zapLogger {
	return &zapLogger{}
}

func (s *zapLogger) ImpleLogger() {}

func (s *zapLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Debug(fmt.Sprintf(format, args...))
}

func (s *zapLogger) Info(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Info(fmt.Sprintf(format, args...))
}

func (s *zapLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Warn(fmt.Sprintf(format, args...))
}

func (s *zapLogger) Error(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Error(fmt.Sprintf(format, args...))
}

// 注意: 此方法将导致程序终止!!!
func (s *zapLogger) Fatal(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Fatal(fmt.Sprintf(format, args...))
}

// 注意: 此方法将导致panic!!!
func (s *zapLogger) Panic(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Panic(fmt.Sprintf(format, args...))
}

// 指定fields
func (s *zapLogger) WithFields(mapFields map[string]interface{}) base.ILogger {
	fields := make([]zap.Field, 0, len(mapFields))
	for key, val := range mapFields {
		fields = append(fields, zap.Any(key, val))
	}

	return &zapFields{fields}
}

func (s *zapLogger) GetLevel() string {
	return _AtomicLevel.Level().String()
}

func GetLogger() *zap.Logger {
	return zap.L()
}
