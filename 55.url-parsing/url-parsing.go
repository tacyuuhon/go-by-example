// URL 提供了一个统一资源定位方式。 这里了解一下 Go 中是如何解析 URL 的。
//
// https://gobyexample.com/url-parsing
package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	// 我们将解析这个 URL 示例，它包含了一个 scheme， 认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 直接访问 scheme。
	fmt.Println("Scheme:", u.Scheme)

	// User 包含了所有的认证信息，这里调用 Username 和 Password 来获取独立值。
	fmt.Println("User:", u.User)
	fmt.Println("User.Username:", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("User.Password:", p)

	// Host 同时包括主机名和端口信息，如过端口存在的话， 使用 strings.Split() 从 Host 中手动提取端口。
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println("Host[0]:", h[0])
	fmt.Println("Host[1]:", h[1])

	// 这里我们提出路径和查询片段信息。
	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)

	// 要得到字符串中的 k=v 这种格式的查询参数，可以使 用 RawQuery 函数。
	// 你也可以将查询参数解析为一个 map。
	// 已解析的查询参数 map 以查询字符串为键，对应 值字符串切片为值，所以如果只想得到一个键对应的第 一个值，将索引位置设置为 [0] 就行了。
	// 运行我们的 URL 解析程序，显示全部我们提取的 URL 的 不同数据块。
	fmt.Println("RawQuery: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("Parse RawQuery:", m)
	fmt.Println("Parse RawQuery[k][0]:", m["k"][0])
}

// Output
/*
	Scheme: postgres
	User: user:pass
	User.Username: user
	User.Password: pass
	host.com:5432
	Host[0]: host.com
	Host[1]: 5432
	Path: /path
	Fragment: f
	RawQuery:  k=v
	Parse RawQuery: map[k:[v]]
	Parse RawQuery[k][0]: v
*/
