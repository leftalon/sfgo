package zap

import (
	"context"
	"fmt"
	"sfgo/core/log/base"

	zapLog "go.uber.org/zap"
)

// 处理带fields的日志
type zapFields struct {
	fields []zapLog.Field
}

func (s *zapFields) ImpleLogger() {}

func (s *zapFields) Debug(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Debug(fmt.Sprintf(format, args...), s.fields...)
}

func (s *zapFields) Info(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Info(fmt.Sprintf(format, args...), s.fields...)
}

func (s *zapFields) Warn(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Warn(fmt.Sprintf(format, args...), s.fields...)
}

func (s *zapFields) Error(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Error(fmt.Sprintf(format, args...), s.fields...)
}

func (s *zapFields) Fatal(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Fatal(fmt.Sprintf(format, args...), s.fields...)
}

func (s *zapFields) Panic(ctx context.Context, format string, args ...interface{}) {
	GetLogger().Panic(fmt.Sprintf(format, args...), s.fields...)
}

// 支持链式调用
func (s *zapFields) WithFields(mapFields map[string]interface{}) base.ILogger {
	if s.fields == nil {
		s.fields = make([]zapLog.Field, 0, len(mapFields))
	}

	for key, val := range mapFields {
		s.fields = append(s.fields, zapLog.Any(key, val))
	}

	return s
}

func (s *zapFields) GetLevel() string {
	return _AtomicLevel.Level().String()
}
