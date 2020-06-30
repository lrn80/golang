package main

import (
	"fmt"
	"math"
)

func varZeroValue()  {
	var a int      // 使用 var 定义
	var str string
	var  d , e float32
	b, c := 3, true // 定义变量的时候赋初始值
	fmt.Println(a, str, d, e,  b, c) // 0  0 0 3 true
}

// 常量的定义
func consts()  {
	const a, b = 3, 4
	const (
		php = 1
		python = 2
		js = "javasrcipt"
	)

	fmt.Println(php, python, js, a, b) // 1 2 javasrcipt 3 4
}

func enums()  {
	const (
		cpp = iota // iota自动加一
		java
		ruby
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(cpp, java, ruby) // 0 1 2
	fmt.Println(b, kb, mb, gb) // 1 1024 1048576 1073741824
}

func main()  {
	fmt.Println("Hello World")  // Hello World

	fmt.Printf("%t\n", 1==2)
	fmt.Printf("二进制：%b\n", 255)
	fmt.Printf("八进制：%o\n", 255)
	fmt.Printf("十六进制：%X\n", 255)
	fmt.Printf("十进制：%d\n", 255)
	fmt.Printf("浮点数：%f\n", math.Pi)
	fmt.Printf("字符串：%s\n", "hello world")

	varZeroValue()
	consts()
	enums()
}
