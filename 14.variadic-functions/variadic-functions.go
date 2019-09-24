// [可变参数函数](https://zh.wikipedia.org/wiki/%E5%8F%AF%E8%AE%8A%E5%8F%83%E6%95%B8%E5%87%BD%E6%95%B8)。
// 在调用时可以用任意数量的参数。 例如，fmt.Println 是一个常见的变参函数。
// Go 函数的另一个关键的方面是闭包结构，这是接下来我们需要看看的。
//
// https://books.mlog.club/gobyexample/variadic-functions.html
package main

import "fmt"

// 这个函数接受任意数目的 int 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 变参函数使用常规的调用方式，传入独立的参数。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果你有一个含有多个值的 slice，想把它们作为参数使用，你要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// Output
/*
	[1 2] 3
	[1 2 3] 6
	[1 2 3 4] 10
*/
