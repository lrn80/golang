package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println(arr) // [0 0 0] 初始值为 0

	arr1 := [3]int{2, 3,4}
	 fmt.Println(arr1) // [2 3 4]

	arr2 := [...]int{5, 6, 7, 8}
	 fmt.Println(arr2) // [5 6 7 8] 未给出固定的初始长度

	var grid [4][5]int
	 fmt.Println(grid) // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]] 二维数组 4行 5列

	// 数组的遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i]); // 5
							  // 6
							  // 7
							  // 8
	}

	// 数组的遍历
	for i, v := range arr2{
		fmt.Println(i, v); // 0 5
		                   // 1 6
		                   // 2 7
		                   // 3 8
	}
}
