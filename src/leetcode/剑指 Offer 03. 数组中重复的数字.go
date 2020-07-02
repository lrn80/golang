package main

import "sort"

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	v := nums[0]
	for _, num := range nums[1 : ]{
		if num == v {
			return num
		} else {
			v = num
		}
	}
	return -1
}
func main() {
	
}
