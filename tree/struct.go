package main

import "fmt"

type Person struct {
	name string
	age int
	email string
}
func main() {
	// 初始化
	person := Person{"Tom", 20, "tom@qq.com"}
	fmt.Println(person) // 输出{Tom 20 tom@qq.com}

	pPerson := &person

	fmt.Println(pPerson) // 输出 &{Tom 20 tom@qq.com}

	pPerson.name = "ruoning"
	pPerson.email = "ruoningli@qq.com"

	fmt.Println(pPerson) // 输出 &{ruoning 20 ruoningli@qq.com}
}
