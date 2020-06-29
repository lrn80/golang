package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	depth = maxDepthRe(root, depth)
	return depth
}

func maxDepthRe(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}

	left := maxDepthRe(root.Left, depth)
	right := maxDepthRe(root.Right, depth)

	if left > right {
		depth = left + 1
 	} else {
 		depth = right + 1
	}

	return depth
}
