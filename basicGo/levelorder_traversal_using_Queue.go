package main

import "fmt"

// Binary Tree Node
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func getTreeNode(data int) *TreeNode {
	// return new TreeNode
	return &TreeNode{
		data,
		nil,
		nil,
	}
}

// Create Q node
type QNode struct {
	data  *TreeNode
	next  *QNode
	level int
}

func getQNode(node *TreeNode) *QNode {
	// return new QNode
	return &QNode{
		node,
		nil,
		0,
	}
}

type MyQueue struct {
	head  *QNode
	tail  *QNode
	count int
}

func getMyQueue() *MyQueue {
	// return new MyQueue
	return &MyQueue{
		nil,
		nil,
		0,
	}
}
func (this MyQueue) size() int {
	return this.count
}
func (this MyQueue) isEmpty() bool {
	return this.count == 0
}

// Add new node of queue
func (this *MyQueue) enqueue(value *TreeNode) {
	// Create a new node
	var node *QNode = getQNode(value)
	if this.head == nil {
		// Add first element into queue
		this.head = node
	} else {
		// Add node at the end using tail
		this.tail.next = node
		// Change node level
		node.level = this.head.level + 1
	}
	this.count++
	this.tail = node
}

// Delete a element into queue
func (this *MyQueue) dequeue() {
	if this.head == nil {
		// Empty Queue
		return
	}
	// Visit next node
	this.head = this.head.next
	this.count--
	if this.head == nil {
		// When deleting a last node of linked list
		this.tail = nil
	}
}

// Get front node
func (this MyQueue) peek() *TreeNode {
	if this.head == nil {
		return nil
	}
	return this.head.data
}

// Level of head node
func (this MyQueue) peekLevel() int {
	if this.size() == 0 {
		return 0
	}
	return this.head.level
}

type BinaryTree struct {
	root *TreeNode
}

func getBinaryTree() *BinaryTree {
	// return new BinaryTree
	return &BinaryTree{
		nil,
	}
}
func (this BinaryTree) levelLineByLine() {
	if this.root != nil {
		var q *MyQueue = getMyQueue()
		// Add first node
		q.enqueue(this.root)
		var node *TreeNode = this.root
		var level int = 0
		for q.isEmpty() == false && node != nil {
			// Important to control new line
			if q.peekLevel() != level {
			}
			if node.left != nil {
				// Add left child node
				q.enqueue(node.left)
			}
			if node.right != nil {
				// Add right child node
				q.enqueue(node.right)
			}
			// Display level node
			fmt.Print("  ", node.data)
			// Change level
			level = q.peekLevel()
			// Remove current node
			q.dequeue()
			// Get new head
			node = q.peek()
		}
	} else {
		fmt.Println("Empty Tree")
	}
}
func main() {
	// Create new tree
	var tree *BinaryTree = getBinaryTree()
	// insert node of binary tree
	tree.root = getTreeNode(1)
	tree.root.left = getTreeNode(2)
	tree.root.right = getTreeNode(3)
	tree.root.left.left = getTreeNode(4)
	tree.root.left.right = getTreeNode(5)
	tree.root.right.left = getTreeNode(6)
	tree.root.right.right = getTreeNode(7)

	// Display level order element
	tree.levelLineByLine()
}

/*
input: 1->left 2, 1->right 3, 2->left4, 2->right5, 3->left 6, 3->right7
Output:-
PS C:\Users\RD\Desktop\GolangCORIOLIS\basicGo> go run levelorder_traversal_using_Queue.go
  1  2  3  4  5  6  7
*/
