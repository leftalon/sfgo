// zap是日志功能的一种实现。
// 支持特性：
// 1. 支持2种日志格式: json 和 logfmt
// 2. 支持日志轮换: 可按文件大小，时间轮换，可设置保留备份文件数。
// 3. 动态调整日志等级:
//
// 如果启动了level server，可以通过 http://localhost:80001/log_level 查看并动态调整日志等级。
//
// 	curl -X PUT -d '{"level":"debug"}' http://localhost:8003/log_level
// 	curl -X GET http://localhost:8003/log_level
package zap

import (
	"fmt"
	"github.com/jsternberg/zap-logfmt"
	zapLog "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	stdlog "log"
	"net/http"
	"os"
	"sfgo/core/log/base"
	"sfgo/core/safego"
	"time"
)

var (
	_LOGHOSTNAME = ""
	_AtomicLevel zapLog.AtomicLevel
)

// TODO log by date

// 初始化，并返回日志对象。
func New() *zapLogger {
	// 初始化全局变量
	_LOGHOSTNAME, _ = os.Hostname()

	// 设置时间戳格式
	encoderConfig := zapLog.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		ts = ts.Local()
		if !base.LogConfigure.LocalTime {
			ts = ts.UTC()
		}
		encoder.AppendString(ts.Format(base.LogConfigure.TimeFormat))
	}

	// 设置日志格式为
	var encoder zapcore.Encoder
	switch base.LogConfigure.Formatter {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zaplogfmt.NewEncoder(encoderConfig)
	}

	// 设置日志rotate
	var writerSyncer zapcore.WriteSyncer
	switch base.LogConfigure.Output {
	case "", "stdout":
		writerSyncer = zapcore.AddSync(os.Stdout)
	case "stderr":
		writerSyncer = zapcore.AddSync(os.Stderr)
	default:
		if base.LogConfigure.MaxSize == 0 && base.LogConfigure.MaxBackups == 0 && base.LogConfigure.MaxAge == 0 {
			// 未启动日志切换
			ws, _, err := zapLog.Open(base.LogConfigure.Output)
			if err != nil {
				stdlog.Fatalf("Failed open log file: %s", base.LogConfigure.Output)
				return nil
			}
			writerSyncer = ws
		} else {
			// 启用日志切换
			hook := lumberjack.Logger{
				Filename:   base.LogConfigure.Output,
				MaxSize:    base.LogConfigure.MaxSize,
				MaxBackups: base.LogConfigure.MaxBackups,
				MaxAge:     base.LogConfigure.MaxAge,
				LocalTime:  base.LogConfigure.LocalTime,
				Compress:   base.LogConfigure.Compress}

			writerSyncer = zapcore.AddSync(&hook)
		}
	}

	// info level
	atomicLevel := zapLog.NewAtomicLevel()
	_AtomicLevel = atomicLevel

	// 使用 zapcore
	core := zapcore.NewCore(
		encoder,
		writerSyncer,
		atomicLevel,
	)

	// 生成logger
	logger := zapLog.New(core)

	//显 caller
	//logger = logger.WithOptions(zap.AddCaller())

	//显示stacktrace
	//logger = logger.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

	// 初始化fields
	fs := make([]zapLog.Field, 0)
	fs = append(fs, zapLog.String("hostname", _LOGHOSTNAME))
	logger = logger.WithOptions(zapLog.Fields(fs...))

	// 设置level
	switch base.LogConfigure.Level {
	case "debug":
		atomicLevel.SetLevel(zapLog.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zapLog.InfoLevel)
	case "warning", "warn":
		atomicLevel.SetLevel(zapLog.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zapLog.ErrorLevel)
	case "fatal":
		atomicLevel.SetLevel(zapLog.FatalLevel)
	case "panic":
		atomicLevel.SetLevel(zapLog.PanicLevel)
	default:
		atomicLevel.SetLevel(zapLog.InfoLevel)
	}

	zapLog.ReplaceGlobals(logger)

	// 动态调整level服务
	if base.LevelServer.Enabled {
		mux := http.NewServeMux()
		mux.Handle("/log_level", atomicLevel)
		safego.Go(func() {
			addr := fmt.Sprintf("%s:%d", base.LevelServer.Host, base.LevelServer.Port)
			http.ListenAndServe(addr, mux)
		})
	}

	return NewLogger()
}
