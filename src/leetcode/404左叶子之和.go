package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	return sumOfLeftLeavesRe(root, &sum)
}

func sumOfLeftLeavesRe(root *TreeNode, sum *int)  int{
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}


	sumOfLeftLeavesRe(root.Left, sum)
	sumOfLeftLeavesRe(root.Right, sum)

	return *sum
}

func main() {
	
}
