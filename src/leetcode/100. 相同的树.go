package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p ==nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if isSameTree(p.Left, q.Left) == false || isSameTree(p.Right, q.Right) == false {
		return false
	}

	return true
}

func main() {
	
}
