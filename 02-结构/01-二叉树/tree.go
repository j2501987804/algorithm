package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{
			Val: val,
		}
	}

	if val >= root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}
	return root
}
