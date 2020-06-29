package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricRe(root.Left, root.Right)
}

func isSymmetricRe(l , r *TreeNode) bool {

	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	if isSymmetricRe(l.Left, r.Right) && isSymmetricRe(l.Right, r.Left) {
		return true
	} else {
		return false
	}
}

func main() {
	
}
