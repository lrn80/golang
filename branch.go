package main

import "fmt"

func eval(a, b int, op string) int {
	var result int
	// switch 语句
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("no operator:" + op)
	}
	return result
}

func oddAndEven(x int)  {
	// if 语句
	if x % 2 == 0 {
		// ...
	}

	// if-else
	if x % 2 == 0 {
		// 偶数
	} else {
		// 奇数
	}

	// if - else - if 多分支
	if x < 0 {
		// 负数
	} else if x == 0 {
		// 零
	} else {
		// 正数
	}
}

// 经典的 for 语句在 go 中没有 while
func PrintOneToTen()  {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 1
	// 死循环相当于 for(;;)
	for  {
		if i > 10 {
			break
		}
		i++
	}
}


func main() {
	
	fmt.Println(eval(4, 4, "+"))
	oddAndEven(4)
	PrintOneToTen()
}
