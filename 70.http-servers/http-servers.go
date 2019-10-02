// 使用net/http包编写一个基本的http服务器很容易。
//
// https://gobyexample.com/http-servers
package main

import (
	"fmt"
	"net/http"
)

// net/http服务器中的一个基本概念是处理程序。
// 处理程序是实现http.Handler接口的对象。
// 编写处理程序的一种常见方法是在具有适当签名的函数上使用http.HandlerFunc适配器。
func hello(w http.ResponseWriter, req *http.Request) {

	// 用作处理程序的函数以http.ResponseWriter和http.ResponseWriter作为参数。
	// 响应编写器用于填充HTTP响应。在这里，我们的简单回答是“hello\n”。
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// 这个处理程序通过读取所有的HTTP请求头并将它们回送到响应体中来做一些更复杂的事情。
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// 我们使用http.HandleFunc便捷功能在服务器路由上注册处理程序。
	// 它在net/http包中设置默认路由器，并以函数作为参数。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最后，我们用端口和处理程序调用ListenAndServe。
	// nil告诉它使用我们刚刚设置的默认路由器。
	http.ListenAndServe(":8080", nil)
}

// Output

// $ curl localhost:8080/hello
/*
	hello
*/

// $ curl localhost:8080/headers
//
//	User-Agent: curl/7.54.0
//	Accept: */*
//
