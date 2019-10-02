// 在前面的例子中，我们了解了生成外部进程的知识，
// 当我们需要访问外部进程时需要这样做，但是有时候，我们只想用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。
// 这时候，我们可以使用经典的 [exec] 方法的 Go 实现。
//
// 当我们运行程序时，它会替换为 ls。
// 注意 Go 并不提供一个经典的 Unix fork 函数。
// 通常这不是个问题，因为运行 Go 协程，生成进程和执行进程覆盖了 fork 的大多数使用场景。
//
// [exec](https://en.wikipedia.org/wiki/Exec_(system_call))
//
// https://gobyexample.com/execing-processes
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 在我们的例子中，我们将执行 ls 命令。
	// Go 需要提供我们需要执行的可执行文件的绝对路径，所以我们将使用 exec.LookPath 来得到它（大概是 /bin/ls）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec 需要的参数是切片的形式的（不是放在一起的一个大字 符串）。
	// 我们给 ls 一些基本的参数。
	// 注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec 同样需要使用环境变量。
	// 这里我们仅提供当前的环境变量。
	env := os.Environ()

	// 这里是 syscall.Exec 调用。
	// 如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。
	// 如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// Output
// $ go run execing-processes.go
/*
	total 8
	drwxr-xr-x   3 tacyuuhon  staff    96B 10  2 15:11 .
	drwxr-xr-x  68 tacyuuhon  staff   2.1K 10  2 14:56 ..
	-rw-r--r--   1 tacyuuhon  staff   1.6K 10  2 15:11 execing-processes.go
*/
