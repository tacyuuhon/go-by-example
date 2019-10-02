// [环境变量]是一个为 [Unix 程序传递配置信息]的普遍方式。
// 让我们来看看如何设置，获取并列举环境变量。
//
// 运行这个程序，显示我们在程序中设置的 FOO 的值，然而没有设置的 BAR 是空的。
// 键的列表是由你的电脑情况而定的。
//
// [环境变量](https://zh.wikipedia.org/wiki/%E7%8E%AF%E5%A2%83%E5%8F%98%E9%87%8F)
// [Unix 程序传递配置信息](https://www.12factor.net/config)
//
// https://gobyexample.com/environment-variables
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 使用 os.Setenv 来设置一个键值对。
	// 使用 os.Getenv 获取一个键对应的值。
	// 如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// 使用 os.Environ 来列出所有环境变量键值对。
	// 这个函数会返回一个 KEY=value 形式的字符串切片。
	// 你可以使用 strings.Split 来得到键和值。
	// 这里我们打印所有的键。
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")

		fmt.Println(pair[0])
		//fmt.Println(pair[0], ":", pair[1])
	}
}

// Output
/*
SSH_AUTH_SOCK
PATH
LOGNAME
XPC_SERVICE_NAME
USER
XPC_FLAGS
VSCODE_NLS_CONFIG
VSCODE_NODE_CACHED_DATA_DIR
VSCODE_LOGS
VSCODE_IPC_HOOK
VSCODE_PID
LESS
LSCOLORS
OLDPWD
PAGER
PWD
SHLVL
ZSH
_
AMD_ENTRYPOINT
PIPE_LOGGING
VERBOSE_LOGGING
VSCODE_IPC_HOOK_EXTHOST
VSCODE_HANDLES_UNCAUGHT_ERRORS
VSCODE_LOG_STACK
APPLICATION_INSIGHTS_NO_DIAGNOSTIC_CHANNEL
GOPATH
GOROOT
VSCODE_PREVENT_FOREIGN_INSPECT
ELECTRON_RUN_AS_NODE
CGO_CFLAGS
FOO
*/
