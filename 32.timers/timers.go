// 我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。
// Go 的内置定时器和打点器特性让这些很容易实现。
// 我们将先学习定时器，然后再学习打点器。
//
// 第一个定时器将在程序开始后 ~2s 失效，但是第二个在它没失效之前就停止了。
//
// https://gobyexample.com/timers
package main

import (
	"fmt"
	"time"
)

func main() {

	// 定时器表示在未来某一时刻的独立事件。
	// 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。
	// 这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(time.Second * 2)

	// <-timer1.C 直到这个定时器的通道 C 明确的发送了定时器失效的值之前，将一直阻塞。
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 如果你需要的仅仅是单纯的等待，你需要使用 time.Sleep。
	// 定时器是有用原因之一就是你可以在定时器失效之前，取消这个定时器。
	// 这是一个例子。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

// Outpup
/*
	Timer 1 expired
	Timer 2 stopped
*/
