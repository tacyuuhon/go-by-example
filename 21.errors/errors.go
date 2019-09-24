// 符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。
// 这与使用异常(exception)的 Java 和 Ruby 以及在 C 语言中有时用到的重载(overloaded)的单返回/错误值有着明显的不同。
// Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用。
// 到 Go 官方博客去看这篇[很棒的文章](https://blog.golang.org/error-handling-and-go) 获取更多关于错误处理的信息。
//
// https://books.mlog.club/gobyexample/errors.html
package main

import (
	"errors"
	"fmt"
)

// 按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New 构造一个使用给定的错误信息的基本 error 值。
		return -1, errors.New("can't work with 42")
	}

	// 返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// 通过实现 Error 方法来自定义 error 类型是可以的。
// 这里使用自定义错误类型来表示上面例子中的参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 在这个例子中，我们使用 &argError 语法来建立一个新的结构体，并提供了 arg 和 prob 这两个字段的值。
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {

	// 下面的两个循环测试了各个返回错误的函数。
	// 注意在 if 行内的错误检查代码，在 Go 中是一个普遍的用法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// 你如果想在程序中使用一个自定义错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// Output
/*
	f1 worked: 10
	f1 failed: can't work with 42
	f2 worked 10
	f2 failed: 42-can't work with it
	42
	can't work with it
*/
