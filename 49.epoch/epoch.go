// 一般程序会有获取 [Unix 时间]的秒数，毫秒数，或者微秒数的需要。
// 来看看如何用 Go 来实现。
//
// [Unix 时间](https://zh.wikipedia.org/wiki/UNIX%E6%97%B6%E9%97%B4)
// https://gobyexample.com/epoch
package main

import (
	"fmt"
	"time"
)

func main() {

	// 分别使用带 Unix 或者 UnixNano 的 time.Now 来获取从自[协调世界时](https://zh.wikipedia.org/wiki/%E5%8D%8F%E8%B0%83%E4%B8%96%E7%95%8C%E6%97%B6)起到现在的秒数或者纳秒数。
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意 UnixMillis 是不存在的，所以要得到毫秒数的话， 你要自己手动的从纳秒转化一下。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 你也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

// Output
/*
	2019-10-01 15:07:40.924723 +0900 JST m=+0.000211599
	1569910060
	1569910060924
	1569910060924723000
	2019-10-01 15:07:40 +0900 JST
	2019-10-01 15:07:40.924723 +0900 JST
*/
