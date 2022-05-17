package main

import (
	"fmt"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}
func printInorder(n *BinaryNode) {
	if n == nil {
		return
	} else {
		printInorder(n.left)
		fmt.Println(n.data)
		printInorder(n.right)
	}
}
func printPreorder(n *BinaryNode) {
	if n == nil {
		return
	} else {
		fmt.Println(n.data)
		printPreorder(n.left)
		printPreorder(n.right)
	}
}

/*
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
*/
func main() {
	tree := &BinaryTree{}
	tree.insert(10).
		insert(20).
		insert(30).
		insert(40).
		insert(50).
		insert(60).
		insert(70).
		insert(80).
		insert(90).
		insert(100).
		insert(110).
		insert(11)
	//print(os.Stdout, tree.root, 0, 'M')

	fmt.Println("Inorder Traversal of Binary Search Tree: ")
	printInorder(tree.root)

	fmt.Println("Preorder Traversal of Binary Search Tree: ")
	printPreorder(tree.root)
}
