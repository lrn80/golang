package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "211"
	c, _ := strconv.Atoi(str[2: 3])
	fmt.Println(c)
}
