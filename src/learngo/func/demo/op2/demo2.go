package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

type calculateFunc func(x , y int) (int, error)

func genCalculator(op operate) calculateFunc {

	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}

		return op(x, y), nil
	}
}

func main() {
	op := func(x, y int) int {
		return x + y
	}

	add := genCalculator(op)
	fmt.Println(add(5,6))

	fmt.Println(calculate(3, 4, op))

}
