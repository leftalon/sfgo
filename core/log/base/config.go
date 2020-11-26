package base

import "sfgo/core/config"

// 日志配置格式

var LogConfigure struct {
	Output     string
	Formatter  string
	Level      string
	LocalTime  bool
	TimeFormat string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var LevelServer struct {
	Enabled bool
	Host    string
	Port    int
}

func init() {
	LogConfigure.Output = "stdout"
	LogConfigure.Level = "info"
	LogConfigure.Formatter = "logfmt"
	LogConfigure.TimeFormat = "2006-01-02T15:04:05.000Z07:00"
	LogConfigure.LocalTime = true
	LogConfigure.MaxSize = 100 // 单位M
	LogConfigure.MaxBackups = 3
	LogConfigure.MaxAge = 30
	LogConfigure.Compress = false

	// 日志级别动态更改
	LevelServer.Enabled = false
	LevelServer.Host = "localhost"
	LevelServer.Port = 80001

	config.Register("sfgo.log", &LogConfigure)
	config.Register("sfgo.log.levelserver", &LevelServer)
}
