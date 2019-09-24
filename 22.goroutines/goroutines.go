// Go 协程(goroutine) 在执行上来说是轻量级的线程。
//
// 当我们运行这个程序时，将首先看到阻塞式调用的输出，然后是 两个 Go 协程的交替输出。
// 这种交替的情况表示 Go 运行时是以并发的方式运行协程的。
//
// 接下来我们将学习在并发的 Go 程序中的 Go 协程的辅助 特性：通道(channels)。
//
// https://gobyexample.com/goroutines
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设我们有一个函数叫做 f(s)。一般会这样同步(synchronously)调用
	f("direct")

	// 使用 go f(s) 在一个 Go 协程中调用这个函数。 这个新的 Go 协程将会并发(concurrently)执行这个函数。
	go f("goroutine")

	// 你也可以为匿名函数启动一个 Go 协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步(asynchronously)运行，所以 程序直接运行到这一行。
	// 这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Scanln()
	fmt.Println("done")
}

// Output
/*
	direct : 0
	direct : 1
	direct : 2
	goroutine : 0
	goroutine : 1
	goroutine : 2
	going

	done
*/
