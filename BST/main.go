package main

import (
	"fmt"
)

func main() {
	bs := NewBST()
	bs = bs.Insert(9)
	bs = bs.Insert(4)
	bs = bs.Insert(6)
	bs = bs.Insert(20)
	bs = bs.Insert(170)
	bs = bs.Insert(15)
	bs = bs.Insert(1)
	bs.PrintBST()
}

type (
	BST struct {
		Root *BinaryTreeNode
	}

	BinaryTreeNode struct {
		LeftNode  *BinaryTreeNode
		RightNode *BinaryTreeNode
		Value     int
	}
)

func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

func (bst *BST) Insert(value int) *BST {
	if bst.Root == nil {
		bst.Root = &BinaryTreeNode{
			Value: value,
		}
		return bst
	}
	current := bst.Root
	for {
		if current.Value >= value {
			if current.LeftNode != nil {
				current = current.LeftNode
			} else {
				// Just insert the node here
				fmt.Printf("JUST current node %d LeftNode: %d\n", current.Value, value)
				current.LeftNode = &BinaryTreeNode{
					Value: value,
				}
				return bst
			}
		} else {
			if current.RightNode != nil {
				current = current.RightNode
			} else {
				fmt.Printf("JUST current node %dRightNode: %d\n", current.Value, value)
				current.RightNode = &BinaryTreeNode{
					Value: value,
				}
				return bst
			}
		}
	}
}

func (bst *BST) PrintBST() {
	Print(bst.Root)
}

func Print(n *BinaryTreeNode) {
	fmt.Printf("    N:%d\n", n.Value)
	if n.LeftNode != nil {
		fmt.Printf("L:%d\n", n.LeftNode.Value)
		Print(n.LeftNode)
	}
	if n.RightNode != nil {
		fmt.Printf("       R:%d\n", n.RightNode.Value)
		Print(n.RightNode)
	}
}
