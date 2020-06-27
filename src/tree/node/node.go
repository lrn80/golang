package node

type TreeNode struct {
	value int
	left, right *TreeNode
}

func (ac *TreeNode) SetValue(value int)  {
	ac.value = value
}

//func main() {
//	var root treeNode
//	root = treeNode{value: 4}
//	fmt.Println(root)
//	root.right.setValue(3)
//	fmt.Println(root.right)
//	//root.left = &treeNode{}
//	//root.right = &treeNode{5, nil, nil}
//	//fmt.Println(root) // 输出 {4 0xc0000044a0 0xc0000044c0}
//	//
//	//root.right.left = new(treeNode) // 建了一个新的treeNode
//	//fmt.Println(root.right) // &{5 0xc000004520 <nil>}
//	//
//	//root.right.right.left = new(treeNode)
//	//fmt.Println(root.right.right)
//}
