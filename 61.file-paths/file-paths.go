// 该filepath软件包提供了以操作系统之间可移植的方式解析和构造文件路径的功能。
// 例如，在Linux: dir/file， 在Windows上: dir\file。
//
// https://gobyexample.com/file-paths
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// Join应该以可移植的方式构造路径。它接受任意数量的参数，并根据它们构造层次结构路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// 您应该始终使用Join而不是手动连接/或\。
	// 除了提供可移植性之外，Join还将通过删除多余的分隔符和目录更改来规范化路径。
	fmt.Println("dir1// join:", filepath.Join("dir1//", "filename"))
	fmt.Println("dir1/../dir1 join:", filepath.Join("dir1/../dir1", "filename"))

	// Dir并且Base可用于拆分目录和文件的路径。
	// 或者，Split将在同一个调用中返回两者。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// 我们可以检查路径是否是绝对的。
	fmt.Println("dir/file is abs:", filepath.IsAbs("dir/file"))
	fmt.Println("/dir/file is abs:", filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// 某些文件名的扩展名后带有一个点。
	// 我们可以使用来将扩展名从此类名称中分离出来Ext。
	ext := filepath.Ext(filename)
	fmt.Println("Ext:", ext)

	// 要查找已删除扩展名的文件名，请使用strings.TrimSuffix。
	fmt.Println("TrimSuffix:", strings.TrimSuffix(filename, ext))

	// Rel查找基准与目标之间的相对路径。如果无法相对于基准目标，则返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("Rel:", rel)
}

// Output
/*
	p: dir1/dir2/filename
	dir1// join: dir1/filename
	dir1/../dir1 join: dir1/filename
	Dir(p): dir1/dir2
	Base(p): filename
	dir/file is abs: false
	/dir/file is abs: true
	Ext: .json
	TrimSuffix: config
	Rel: t/file
*/
