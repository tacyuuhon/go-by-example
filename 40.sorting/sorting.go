// Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。
// 我们首先关注内置数据类型的排序。
//
// https://gobyexample.com/sorting
package main

import (
	"fmt"
	"sort"
)

func main() {

	// 排序方法是针对内置数据类型的；这里是一个字符串的例子。
	// 注意排序是原地更新的，所以他会改变给定的序列并且不返回 一个新值。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个 int 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 我们也可以使用 sort 来检查一个序列是不是已经是排好序的。
	// 运行程序，打印排序好的字符串和整型序列以及我们 AreSorted 测试的结果 true。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

// Output
/*
	Strings: [a b c]
	Ints: [2 4 7]
	Sorted:  true
*/
