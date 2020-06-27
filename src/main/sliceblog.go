package main

import "fmt"

func main() {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr) // 输出 [0 1 2 3 4 5]

	slice := arr[1:4]
	fmt.Println(slice) // 输出 [1 2 3] arr[1]到arr[3]不包括arr[4]

	slice[1] = 7
	fmt.Println(arr)   // 输出 [0 1 7 3 4 5]
	fmt.Println(slice) // 输出 [1 7 3]

	s1 := make([]int, 3)
	fmt.Printf("The len of s1: %d\n", len(s1))  // 输出 The len of s1: 3
	fmt.Printf("The cap of s1: %d\n", cap(s1))  // 输出 The cap of s1: 3
	fmt.Printf("The value of s1: %d\n", s1)  // 输出 The value of s1: [0 0 0]

	s1 = append(s1, 1, 2)

	fmt.Printf("The len of s1: %d\n", len(s1))  // 输出 The len of s1: 5
	fmt.Printf("The cap of s1: %d\n", cap(s1))  // 输出 The cap of s1: 6
	fmt.Printf("The value of s1: %d\n", s1)  // 输出 The value of s1: [0 0 0 1 2]

	s3 := make([]int, 1024)

	fmt.Printf("The len of s1: %d\n", len(s3))  // 输出 The len of s1: 1024
	fmt.Printf("The cap of s1: %d\n", cap(s3))  // 输出 The cap of s1: 1024


	s3 = append(s3, 1, 2)

	fmt.Printf("The len of s1: %d\n", len(s3))  // 输出 The len of s1: 1026
	fmt.Printf("The cap of s1: %d\n", cap(s3))  // 输出 The cap of s1: 1280


	s2 := make([]int, 3, 6)
	fmt.Printf("The len of s2: %d\n", len(s2)) // 输出 The len of s2: 3
	fmt.Printf("The cap of s2: %d\n", cap(s2)) // 输出 The cap of s2: 6
	fmt.Printf("The value of s2: %d\n", s2)  // 输出 The value of s2: [0 0 0]

	arr1 := [...]int{0, 1, 2, 3, 4}
	s4 := arr1[0:5] // [0,5)
	s5 := s4
	s5 = append(s5, 6)
	s4[1] = 10
	fmt.Println(s4) // 输出 [0 10 2 3 4]
	fmt.Println(s5) // 输出 [0 1 2 3 4 6]
}
