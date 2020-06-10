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

func main()  {
	fmt.Println("Hello World")
	varZeroValue()
}
