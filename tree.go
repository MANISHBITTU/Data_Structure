package main

import (
	"fmt"
	"math"
)

func main() {
	tree := new(Tree)
	tree.insert(10)
	tree.insert(15)
	tree.insert(17)
	tree.insert(13)
	tree.insert(5)
	tree.insert(3)
	tree.insert(7)
	//	tree.root.printInorder()
	//	tree.findMin()
	fmt.Print(tree.root.totalLeafNode(0))
}

type Node struct {
	data        int
	left, right *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) insert(val int) {
	if tree.root == nil {
		tree.root = &Node{val, nil, nil}
		return
	}
	tree.root.insert(val)
}
func (node *Node) insert(val int) {
	if node == nil {
		return
	} else {
		if val <= node.data {
			if node.left == nil {
				node.left = &Node{val, nil, nil}
			} else {
				node.left.insert(val)
			}
		} else {
			if node.right == nil {
				node.right = &Node{val, nil, nil}
			} else {
				node.right.insert(val)
			}
		}
	}

}
func (node *Node) printInorder() {
	if node != nil {
		node.left.printInorder()
		fmt.Print(node.data, " ")
		node.right.printInorder()
	}
	return
}
func (tree *Tree) findMin() {
	node := tree.root
	if node == nil {
		fmt.Println("Tree is empty")
	} else {
		for node.left != nil {
			node = node.left
		}
		fmt.Println(node.data, " is minimum value")
	}
}

func (tree *Tree) IsBst() {
	result := isBst(tree.root, math.MinInt32, math.MaxInt32)
	if result {
		fmt.Println("Tree os BST")
	} else {
		fmt.Println("Tree is not BST")
	}
}

func isBst(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.data < min || node.data > max {
		return false
	}
	return isBst(node.left, min, node.data) && isBst(node.right, node.data, max)

}

func (node *Node) height() int {
	if node == nil {
		return 0
	}
	ldefth := node.left.height() + 1
	rdefth := node.right.height() + 1
	if ldefth > rdefth {
		return ldefth
	} else {
		return rdefth
	}
}

func (node *Node) totalLeafNode(count int) int {
	if node == nil {
		return 0
	} else {
		if node.left == nil && node.right == nil {
			count = count + 1
		}
		if node.left != nil {
			count = node.left.totalLeafNode(count)
		}
		if node.right != nil {
			count = node.right.totalLeafNode(count)
		}
		return count
	}

}

fun printAllPaths(node *Node,stk *Stack){
	if(node==nil){
		return
	}
	stk.push(node.data)
	if(node.left==nil&&node.right==nil){
		stk.print()
		stk.pop()
		return
	}
	printAllPaths(node.left,stk)
	printAllPaths(node.right,stk)
}
