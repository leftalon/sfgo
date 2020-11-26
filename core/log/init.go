package log

import (
	"sfgo/core/log/zap"
)

func init() {
	DefaultLogger = zap.New()
}
