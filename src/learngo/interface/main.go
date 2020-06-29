package main

import (
	"fmt"
	real2 "learngo/interface/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("https://liruoning.cn")
}

func main() {
	var r Retriever  // 静态类型
	// 动态类型
	r = real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Println(download(r))
	fmt.Printf("%T %v\n", r, r)
}
