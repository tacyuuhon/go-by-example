// 在文件系统中，GO有几个有用的函数。
//
// https://gobyexample.com/directories
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

	// 在当前工作目录中创建新的子目录。
	err := os.Mkdir("subdir", 0755)
	check(err)

	// 创建临时目录时，最好推迟删除它们。RemoveAll将删除整个目录树（类似于rm-rf）。
	defer os.RemoveAll("subdir")

	// helper函数创建新的空文件。
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// 我们可以创建目录层次结构，包括使用MkdirAll的父级。这类似于命令行mkdir-p。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir列出目录内容，返回os.FileInfo对象的一部分。
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir允许我们更改当前的工作目录，类似于cd。
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 现在我们将在列出当前目录时看到subdir/parent/child的内容。
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd回到我们开始的地方。
	err = os.Chdir("../../..")
	check(err)

	// 我们也可以递归地访问目录，包括它的所有子目录。
	// Walk接受一个回调函数来处理访问的每个文件或目录。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// 对于filepath.Walk递归找到的每个文件或目录，都会调用一次visit。
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}

// Output
/*
	Listing subdir/parent
	child true
	file2 false
	file3 false
	Listing subdir/parent/child
	file4 false
	Visiting subdir
	subdir true
	subdir/file1 false
	subdir/parent true
	subdir/parent/child true
	subdir/parent/child/file4 false
	subdir/parent/file2 false
	subdir/parent/file3 false
*/
