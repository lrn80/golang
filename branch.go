package main

import "io/ioutil"
import "fmt"

func eval(a, b int, op string) int {
	var result int
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

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(eval(4, 4, "+"))


}
