package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 3
	ch1 <- 1
	ch1 <- 4
	fmt.Printf("The first element received from channel ch1: %v\n", <-ch1)
}
