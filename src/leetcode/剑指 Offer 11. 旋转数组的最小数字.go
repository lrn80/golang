package main

import "math"

func minArray(numbers []int) int {
	if len(numbers) < 1 {
		return -1
	}

	pre := 0
	next := len(numbers) - 1
	mid := pre

	for numbers[pre] >= numbers[next] {

		if (next - pre) < 1 {
			mid = next
			break
		}

		mid = int(math.Floor(float64(next + pre) / 2))

		if numbers[mid] == numbers[pre] && numbers[mid] == numbers[next] {
			return MinInOrder(numbers, pre, next)
		}

		if numbers[pre] <= numbers[mid] {
			pre = mid
		}

		if numbers[next] >= numbers[mid] {
			next = mid
		}

	}

	return numbers[mid]
}

func MinInOrder(numbers []int, pre, next int) int {
	res := numbers[pre]
	for i := pre +1; i <= next; i++ {
		if res > numbers[i] {
			res = numbers[i]
		}
	}
	return res
}

func main() {
	
}
