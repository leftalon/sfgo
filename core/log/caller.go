package log

import (
	"fmt"
	"runtime"
	"strings"
)

func Caller() map[string]interface{} {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return map[string]interface{}{}
	}

	_func := runtime.FuncForPC(pc)
	if _func == nil {
		return map[string]interface{}{}
	}
	function := _func.Name()

	parts := strings.Split(file, "/")
	src := parts[len(parts)-1]

	return map[string]interface{}{"caller": fmt.Sprintf("%s(%s:%d)", function, src, line)}
}
