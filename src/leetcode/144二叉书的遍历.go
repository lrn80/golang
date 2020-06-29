package main

import node "leetcode/TreeNode"

func preorderTraversal(root *node.TreeNode) []int {
	res := []int{}
	preorderRe(root, &res)
	return res
}

func preorderRe(root *node.TreeNode, res *[]int)  {

	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderRe(root.Left, res)
	preorderRe(root.Right, res)

}


