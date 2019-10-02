// 单元测试是编写原则性GO程序的一个重要部分。
// 测试包提供了编写单元测试所需的工具，并且使用 go test 命令运行测试。
// 为了演示起见，此代码位于包main中，但它可以是任何包。
// 测试代码通常与测试的代码在同一个包中。
//
// https://gobyexample.com/testing
package main

import (
	"fmt"
	"testing"
)

// 我们将测试一个整数最小值的简单实现。
// 通常，我们正在测试的代码将位于名为intutils.go之类的源文件中，然后将其测试文件命名为intutils_test.go。
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b

}

// 通过编写名称以Test开头的函数来创建测试。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		// t.Error* 将报告测试失败，但继续执行测试。t.Fail* 将报告测试失败并立即停止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 编写测试可能是重复的，因此使用表驱动样式是惯用的，
// 其中测试输入和预期输出列在表中，单个循环遍历它们并执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run允许运行"子测试"，每个表条目一个。在执行go test-v时，这些将单独显示。
	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Output
// 以详细模式运行当前项目中的所有测试。
// $ go test -v
/*
	=== RUN   TestIntMinBasic
	--- PASS: TestIntMinBasic (0.00s)
	=== RUN   TestIntMinTableDriven
	=== RUN   TestIntMinTableDriven/0,1
	=== RUN   TestIntMinTableDriven/1,0
	=== RUN   TestIntMinTableDriven/2,-2
	=== RUN   TestIntMinTableDriven/0,-1
	=== RUN   TestIntMinTableDriven/-1,0
	--- PASS: TestIntMinTableDriven (0.00s)
		--- PASS: TestIntMinTableDriven/0,1 (0.00s)
		--- PASS: TestIntMinTableDriven/1,0 (0.00s)
		--- PASS: TestIntMinTableDriven/2,-2 (0.00s)
		--- PASS: TestIntMinTableDriven/0,-1 (0.00s)
		--- PASS: TestIntMinTableDriven/-1,0 (0.00s)
	PASS
	ok      github.com/tacyuuhon/go-by-example/64.testing   0.006s
*/
