// SHA1 散列经常用生成二进制文件或者文本块的短标识。
// 例如，git 版本控制系统大量的使用 SHA1 来标识受版本控制的文件和目录。
// 这里是 Go 中如何进行 SHA1 散列计算的例子。
// Go 在多个 crypto/* 包中实现了一系列散列函数。
// 如果你需要密码学上的安全散列，你需要小心的研究一下[哈希强度]。
//
// [哈希强度](https://en.wikipedia.org/wiki/Cryptographic_hash_function)
//
// https://gobyexample.com/sha1-hashes
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)， 然后 sha1.Sum([]byte{})。
	// 这里我们从一个新的散列开始。
	s := "sha1 this string"
	h := sha1.New()

	// 写入要处理的字节。
	// 如果是一个字符串，需要使用 []byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))

	// 这个用来得到最终的散列值的字符切片。
	// Sum 的参数可以用来给现有的字符切片追加额外的字节切片：一般不需要。
	bs := h.Sum(nil)

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。
	// 使用 %x 来将散列结果格式化为 16 进制字符串。
	// 运行程序计算散列值并以可读 16 进制格式输出。
	// 你可以使用和上面相似的方式来计算其他形式的散列值。
	// 例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

// Outout
/*
	sha1 this string
	cf23df2207d99a74fbe169e3eba035e633b65d94
*/
