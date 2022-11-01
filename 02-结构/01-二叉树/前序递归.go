/*
前序遍历：先打印根节点，然后递归打印左边，再右边
*/
package main

func preorderTraversal(root *Node)  {
	if root == nil {
		return
	}
	println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}
