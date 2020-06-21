package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:5] // [2,5)
	fmt.Println(s) // [2 3 4]

	s = arr[:9] // 从 a[0] 到 a[9],但是不包括 a[9]
	fmt.Println(s)

	s = arr[2:] // 从a[2] 到 a[9]，且包括 a[2]
	fmt.Println(s)
}
