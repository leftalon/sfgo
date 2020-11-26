package base

import "context"

type ILogger interface {
	ImpleLogger()
	Debug(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, format string, args ...interface{})
	Fatal(ctx context.Context, format string, args ...interface{})
	Panic(ctx context.Context, format string, args ...interface{})
	WithFields(mapFields map[string]interface{}) ILogger
	GetLevel() string
}
