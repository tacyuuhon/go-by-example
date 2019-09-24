// 在 Go 中，变量被显式声明，并可以被编译器用来检查函数调用时的类型正确性。
//
// https://gobyexample.com/variables
package main

import "fmt"

func main() {
	// var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	// 你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 将自动推断已经初始化的变量类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的初始化值是 0。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写，例如这个例子中的 var f string = "short"。
	f := "short"
	fmt.Println(f)
}

// Output
/*
	initial
	1 2
	true
	0
	short
*/
