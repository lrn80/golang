package main

import "fmt"

// 返回一个值
func max(a, b int)int  {
	if a > b {
		return a
	} else {
		return b
	}
}

// 返回多个值
func swap(a, b int) (int, int) {
	return b, a
}

// 不定参数
func sum(nums ...int)  {
	fmt.Print(nums, " ")  //输出如 [1, 2, 3] 之类的 数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}
func main() {
	a, b := swap(3, 4)
	fmt.Println(a, b) // 4 3
	fmt.Println(max(a, b)) // 4

	sum(1, 2)
	sum(1, 2, 3)

	//传数组
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
