package main

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)

	return left + right + 1
}



func main() {
	
}
