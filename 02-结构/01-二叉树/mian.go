package main

func main() {
	// 创建树
	root := Insert(nil,8)
	root = Insert(root,7)
	root = Insert(root,3)
	root = Insert(root,20)
	root = Insert(root,4)
	root = Insert(root,1)
	root = Insert(root,2)

	// 前序
	preorderTraversal(root)
}
