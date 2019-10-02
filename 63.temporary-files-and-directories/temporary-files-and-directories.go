// 在整个程序执行过程中，我们通常希望创建程序退出后不需要的数据。
// 临时文件和目录对于这个目的很有用，因为它们不会随着时间污染文件系统。
//
// https://gobyexample.com/temporary-files-and-directories
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 创建临时文件的最简单方法是调用ioutil.TempFile。
	// 它创建一个文件并打开它进行读写。
	// 我们提供""作为第一个参数，因此在我们操作系统的默认位置中会创建这个文件。
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// 显示临时文件的名称。
	// 在基于Unix的操作系统上，目录可能是/tmp。
	// 文件名以作为ioutil.TempFile的第二个参数的前缀开头，其余部分将自动选择，以确保并发调用始终创建不同的文件名。
	fmt.Println("Temp file name:", f.Name())

	// 做完后把文件整理一下。
	// OS很可能在一段时间后自行清理临时文件，但这是很好的做法。
	defer os.Remove(f.Name())

	// 我们可以把一些数据写入文件。
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 如果我们打算写很多临时文件，我们可能更喜欢创建一个临时目录。
	// ioutil.TempDir的参数与TempFile的相同，但是它返回一个目录名而不是一个打开的文件。
	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// 现在，我们可以通过在临时文件名前面加上临时目录来合成它们。
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

// Output
/*
	Temp file name: /var/folders/h2/tvnwqn0n4zjg2g28k80sk6pm0000gp/T/sample033722911
	Temp dir name: /var/folders/h2/tvnwqn0n4zjg2g28k80sk6pm0000gp/T/sampledir088778994
*/
