// ###安全执行goroutine的模块。
// 用于避免goroutine中panic导致程序退出。
package safego

import (
	"fmt"
	"runtime/debug"

	"go.uber.org/zap"
)

// 用此函数启动goroutine。
//
// 分2种情况：
// 1. goroutine的函数不带参数，可以直接将函数名作为参数启动。
// 2. goroutine的函数需要参数，可以在函数使用闭包外部的变量。此时，只能通过匿名函数方式启动。
func Go(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				zap.L().Error(fmt.Sprintf("Panic recovered: %s, stack: %s", r, string(debug.Stack())))
			}
		}()

		f()
	}()
}

// 较少用。用此函数启动带任意参数的goroutine，参数类型只能是interface{}，在函数内部再进行类型转换。
func GoArgs(f func(...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				zap.L().Error(fmt.Sprintf("Panic recovered: %s, stack: %s", r, string(debug.Stack())))
			}
		}()

		f(args...)
	}()
}
