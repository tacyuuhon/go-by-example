// 我们的第一个程序将打印传说中的 “hello world“ 消息，下边是完整的程序代码。
//
// 要运行这个程序，将代码放到 hello-world.go 中 然后执行 go run 。
// 有时候我们想将程序编译成二进制文件， 可以通过 go build 来达到目的。
// 然后我们可以直接运行这个二进制文件。
// 现在我们可以运行并编译基本的 Go 程序，让我们开始学习更多关于这门语言的知识吧。
//
// https://gobyexample.com/hello-world
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

// Output
/*
	hello world
*/
