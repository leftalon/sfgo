package e_test

import (
	"fmt"
	"sfgo/core/e"
)

func ExampleNew() {
	err := e.New("错误消息")
	fmt.Printf("%s\n", err)
	// Output: 错误消息
}

func ExampleErrorf() {
	err := e.New("错误消息, 参数1: %s, 参数2: %d", "a", 2)
	fmt.Printf("%s\n", err)
	// Output: 错误消息, 参数1: a, 参数2: 2
}

func ExampleWrap() {
	err1 := e.New("原始错误")
	err2 := e.Wrap(err1, "包装后错误消息")
	fmt.Printf("%s\n", err2)
	// Output: 包装后错误消息: 原始错误
}

func ExampleWrap_args() {
	err1 := e.New("原始错误")
	err2 := e.Wrap(err1, "包装后错误消息，参数1: %s, 参数2: %d", "a", 2)
	fmt.Printf("%s\n", err2)
	// Output: 包装后错误消息，参数1: a, 参数2: 2: 原始错误
}
