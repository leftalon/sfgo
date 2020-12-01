/*
### 错误处理模块
* 生成错误：

在将错误值返回上级函数时，包含最原始的错误上下文信息，比如：错误链，错误发生时的stacktrace。

有2个方法：

 1. New：生成全新的错误值。
 2. Wrap：新错误值中封装了原错误值。在外层进行错误处理时，可以追溯到内部任意层的原始错误。

* 错误处理：

*/
package e

import (
	"fmt"
)

// New formats according to a format specifier and returns the string
// as a value that satisfies error.
// New also records the stack trace at the point it was called.
func New(format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrap(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}

	// 只保留最初的一次stacktrace
	//if hasStack(err) {
	//	return err
	//}

	return &withStack{
		err,
		callers(),
	}
}

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
