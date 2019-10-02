// 该encoding.xml软件包提供了对XML和类似XML格式的内置支持。
//
// https://gobyexample.com/xml
package main

import (
	"encoding/xml"
	"fmt"
)

// Plant struct
// 此类型将映射到XML。与JSON示例类似，字段标记包含用于编码器和解码器的指令。
// 在这里，我们使用XML包的一些特殊功能：XMLName字段名称指示表示此结构的XML元素的名称；
//  id,attr表示该Id字段是XML 属性，而不是嵌套元素。
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.ID, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{ID: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 发出代表我们工厂的XML；使用 MarshalIndent以产生更人类可读输出。
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println("---------- XML MarshalIndent Plant ----------")
	fmt.Println(string(out))
	fmt.Println()

	fmt.Println("---------- XML Header Plant ----------")
	// 要将通用XML标头添加到输出，请显式附加。
	fmt.Println(xml.Header + string(out))
	fmt.Println()

	var p Plant
	// 用于Unmarhshal将带有XML的字节流解析为数据结构。
	// 如果XML格式不正确或无法映射到Plant，将返回描述性错误。
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println("---------- XML Unmarshal Plant ----------")
	fmt.Println(p)
	fmt.Println()

	tomato := &Plant{ID: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// 该parent>child>plant字段标签告诉编码器嵌套所有plant为<parent><child>...
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")

	fmt.Println("---------- XML MarshalIndent Nesting ----------")
	fmt.Println(string(out))
}

// Output
/*
	---------- XML MarshalIndent Plant ----------
	<plant id="27">
		<name>Coffee</name>
		<origin>Ethiopia</origin>
		<origin>Brazil</origin>
	</plant>

	---------- XML Header Plant ----------
	<?xml version="1.0" encoding="UTF-8"?>
	<plant id="27">
		<name>Coffee</name>
		<origin>Ethiopia</origin>
		<origin>Brazil</origin>
	</plant>

	---------- XML Unmarshal Plant ----------
	Plant id=27, name=Coffee, origin=[Ethiopia Brazil]

	---------- XML MarshalIndent Nesting ----------
	<nesting>
		<parent>
			<child>
			<plant id="27">
				<name>Coffee</name>
				<origin>Ethiopia</origin>
				<origin>Brazil</origin>
			</plant>
			<plant id="81">
				<name>Tomato</name>
				<origin>Mexico</origin>
				<origin>California</origin>
			</plant>
			</child>
		</parent>
	</nesting>
*/
