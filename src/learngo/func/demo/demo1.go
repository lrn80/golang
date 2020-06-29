package main

import "fmt"

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error)  {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	p = printToStd
	p("something")

	op := func(x, y int) int {
		return x + y
	}

	fmt.Println(calculate(3, 4, op))
}
