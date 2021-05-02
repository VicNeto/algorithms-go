package dataStructures

import "fmt"

type SimpleBinaryTreeNode struct {
	Key   int
	Left  *SimpleBinaryTreeNode
	Right *SimpleBinaryTreeNode
}

func AddNode(t *SimpleBinaryTreeNode, v int) {
	if v < t.Key {
		if t.Left != nil {
			AddNode(t.Left, v)
		} else {
			t.Left = &SimpleBinaryTreeNode{v, nil, nil}
		}
	} else {
		if t.Right != nil {
			AddNode(t.Right, v)
		} else {
			t.Right = &SimpleBinaryTreeNode{v, nil, nil}
		}
	}
}

func PreOrder(t *SimpleBinaryTreeNode) {
	if t == nil {
		return
	}
	fmt.Printf("%d ", t.Key)
	PreOrder(t.Left)
	PreOrder(t.Right)
}

func Order(t *SimpleBinaryTreeNode) {
	if t == nil {
		return
	}
	Order(t.Left)
	fmt.Printf("%d ", t.Key)
	Order(t.Right)
}

func PostOrder(t *SimpleBinaryTreeNode) {
	if t == nil {
		return
	}
	PostOrder(t.Left)
	PostOrder(t.Right)
	fmt.Printf("%d ", t.Key)
}
