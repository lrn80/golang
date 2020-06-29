package main

import "math"

func isBalanced(root *TreeNode) bool {
	depth := 0
	isBalancedRe(root, &depth)

}

func isBalancedRe(root *TreeNode, depth *int) bool{
	left := 0
	right := 0
	if root == nil {
		*depth = 0
		return true
	}

	if isBalancedRe(root.Left, &left) && isBalancedRe(root.Right, &right) {
		if math.Abs(float64(left - right)) <= 1 {
			*depth = int(math.Max(float64(left), float64(right)))
			return true
		} else {
			return false
		}
	}

	return false
}

func main() {
	
}
