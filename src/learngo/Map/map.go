package main

import "fmt"

func main() {

	//Map 定义
	m := map[string]string {
		"name": "ruoning",
		"lastname": "蓝鲸",
		"site": "liruoning.cn",
	}

	// 开始声明的country 为nil 不可以进行赋值
	var country map[string]string
	// 再使用make创建一个非nil的map
	country = make(map[string]string)

	country["beijing"] = "北京"
	country["shanghai"] = "上海"
	country["hangzhou"] = "杭州"

	v := country["beijing"] // 取值

	// 遍历Map
	for k, v := range country{
		// 遍历country 输出的顺序可能不一样
		fmt.Println(k, v) // 输出 beijing 北京
						  // 	 shanghai 上海
						  //     hangzhou 杭州
	}

	fmt.Println(m, country) // 输出 map[lastname:蓝鲸 name:ruoning site:liruoning.cn] map[beijing:北京 hangzhou:杭州 shanghai:上海]


	k := "beijing"
	// 判断这个key值是否存在 存在 ok true 不存在 false
	v, ok := country["beijing"]

	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v) // 输出 The element of key "beijing": %!d(string=北京)
	} else {
		fmt.Println("Not found!")
	}
}
