package main

import (
	"fmt"
	"strconv"
)

var letterMap = []string{
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

var res = []string{}

func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) == 0 {
		return res
	}

	findLetter(digits, "", 0)

	return res
}

func findLetter(digits , s string, index int)  {

	if index == len(digits) {
		res = append(res, s)
		return
	}

	c, _ := strconv.Atoi(digits[index: index + 1])

	letter := letterMap[c - 2]

	for i := 0; i < len(letter); i++ {
		findLetter(digits, s + letter[i: i + 1], index + 1)
	}

	return
}

func main() {
	fmt.Println(letterCombinations("23"))
}
