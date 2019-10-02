// [命令行参数](https://en.wikipedia.org/wiki/Command-line_interface#Arguments) 是指定程序运行参数的一个常见方式。
// 例如，go run hello.go， 程序 go 使用了 run 和 hello.go 两个参数。
// 要实验命令行参数，最好先使用 go build 编译一个可执行 二进制文件
// $ go build command-line-arguments.go
// $ ./command-line-arguments a b c d
//
// https://gobyexample.com/command-line-arguments
package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args 提供原始命令行参数访问功能。
	// 注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 你可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]

	fmt.Println("argsWithProg:", argsWithProg)
	fmt.Println("argsWithoutProg:", argsWithoutProg)
	fmt.Println("arg:", arg)
}

// Output
// go build command-line-arguments.go
// ./command-line-arguments a b c d
/*
	argsWithProg: [./command-line-arguments a b c d]
	argsWithoutProg: [a b c d]
	arg: c
*/
