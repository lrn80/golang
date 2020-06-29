package main

import node "leetcode/TreeNode"

func invertTree(root *node.TreeNode) *node.TreeNode {
	// 循环终止条件
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = swap(root.Left, root.Right)

	return root
}

func swap(lNode *node.TreeNode, rNode *node.TreeNode) (*node.TreeNode, *node.TreeNode) {
	return rNode, lNode
}

