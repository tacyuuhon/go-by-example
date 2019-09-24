// if 和 else 分支结构在 Go 中非常直接。
//
// 这里是一个基本的例子。
// 你可以不要 else 只用 if 语句。
// 在条件语句之前可以有一个声明语句；在这里声明的变量 可以在所有的条件分支中使用。
//
// 注意，在 Go 中条件语句不需要圆括号，但是需要加上花括号。
// Go 里没有三目运算符， 所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句
//
// https://books.mlog.club/gobyexample/if-else.html
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
