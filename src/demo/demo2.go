package main

import (
	"flag"
	"fmt"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main() {
	// 把用户传递的命令行参数解析为对应的值
	// 通过 flag.Parse() 函数接下命令行参数，解析函数将会在碰到第一个非 flag 命令行参数时停止：
	flag.Parse()
	fmt.Printf("hello, %s\n", name)
}
