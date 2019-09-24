// 在 Go 中，变量 被显式声明，并可以被编译器用来 检查函数调用时的类型正确性。
//
// var 声明 1 个或者多个变量。
// 你可以一次性声明多个变量。
// Go 将自动推断已经初始化的变量类型。
// 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。例如，一个 int 的零值是 0。
// := 语法是声明并初始化变量的简写，例如 这个例子中的 var f string = "short"。
//
// https://books.mlog.club/gobyexample/variables.html
package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

}
