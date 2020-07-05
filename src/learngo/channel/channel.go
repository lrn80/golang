package main

import (
	"fmt"
)

func chanDemo()  {
	c := make(chan int)
	go func() {
		for  {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 2
	c <- 1
	//time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
