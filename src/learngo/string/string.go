package main

import (
	"fmt"
)

func main() {
	str := "I love you!"

	bytes := []byte(str) // 转换成字节

	fmt.Println(bytes)  // [73 32 108 111 118 101 32 121 111 117 33]
}
