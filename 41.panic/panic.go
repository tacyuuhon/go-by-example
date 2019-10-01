// panic 意味着有些出乎意料的错误发生。
// 通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
//
// https://gobyexample.com/panic
package main

import "os"

func main() {

	// 我们将在这个网站中使用 panic 来检查预期外的错误。
	// 这个是唯一一个为 panic 准备的例子。
	panic("a problem")

	// panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道（或者不想）处理时终止运行。
	// 这里是一个在创建一个新文件时返回异常错误时的 panic 用法。
	//
	// 运行程序将会引起 panic，输出一个错误消息和 Go 运行时栈信息，并且返回一个非零的状态码。
	// 注意，不像在有些语言中使用异常处理错误，在 Go 中则习惯通过返回值来标示错误。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// Output
/*
	panic: a problem

	goroutine 1 [running]:
	main.main()
		/Users/tacyuuhon/go/src/github.com/tacyuuhon/go-by-example/41.panic/panic.go:13 +0x39
*/
