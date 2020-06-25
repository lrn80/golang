package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{value: 4}
	fmt.Println(root)


}
