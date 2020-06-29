package main

import "fmt"


func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return []string{}
	}
	binaryTreePathsRe(root, "", &res)
	return res
}
func binaryTreePathsRe(root *TreeNode, path string, res *[]string)()  {
	if root == nil {
		return
	}

	if path != "" {
		path = path + "->"
	}
	path = path + fmt.Sprintf("%d", root.Val)
	
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}

	binaryTreePathsRe(root.Left, path, res)
	binaryTreePathsRe(root.Right, path, res)
}
func main() {

}
