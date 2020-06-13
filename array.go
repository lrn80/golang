package main

import "fmt"

func main() {
	var arr [3]int
	 arr1 := [3]int{2, 3,4}
	 arr2 := [...]int{5, 6, 7, 8}
	var grid [4][5]int
	 fmt.Println(arr, arr1, arr2, grid)

	 // 数组的遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i]);
	}
	// 数组的遍历
	for i, v := range arr2{
		fmt.Println(i, v);
	}
}
