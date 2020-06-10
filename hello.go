package main

import "fmt"

func varZeroValue()  {
	var a int      // 使用 var 定义
	var str string
	var  d , e float32
	b, c := 3, true // 使用 : 定义
	fmt.Printf(str)
	fmt.Println(a, b, c, d, e)
}

func consts()  {
	const a, b = 3, 4
	const (
		php = 1
		python = 2
		js = "javasrcipt"
	)

	println(php, python, js, a, b)
}

func enums()  {
	const (
		cpp = iota
		java
		ruby
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	println(cpp, java, ruby)
	println(b, kb, mb, gb)
}

func main()  {
	fmt.Println("Hello World")
	varZeroValue()
	consts()
	enums()
}
