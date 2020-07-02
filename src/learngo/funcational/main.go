package main

import (
	"bufio"
	"fmt"
	"learngo/funcational/fib"
	"os"
)

func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	f := fib.Fib()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("./fib.txt")

	f := fib.Fib()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
