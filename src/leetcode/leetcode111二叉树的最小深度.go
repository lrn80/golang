package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	// root为叶子节点，则直接返回
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}

	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}


