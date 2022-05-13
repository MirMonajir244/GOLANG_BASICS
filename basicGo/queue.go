package main

import (
	"fmt"
	"os"
)

type node struct {
	val  int
	next *node
}
type queue struct {
	front *node
	rear  *node
	len   int
}

func (n *queue) enque(num int) {
	var temp = &node{}
	temp.val = num
	temp.next = nil
	if n.front == nil && n.rear == nil {
		n.front = temp
		n.rear = temp
	} else {
		n.rear.next = temp
		n.rear = temp
	}

	n.len++

}

func (n *queue) deque() {
	if n.front == n.rear {
		n.front = nil
		n.rear = nil
	} else {
		n.front = n.front.next
	}
	n.len--
}

func (n *queue) frontt() {
	fmt.Println("Front of the queue is: ", n.front.val)
}
func (n *queue) size() {
	fmt.Println("Size of the queue is: ", n.len)
}
func (n *queue) display() {
	var p *node = n.front
	for p.next != nil {
		fmt.Printf("%d <-", p.val)
		p = p.next
	}
}
func main() {
	var myQ = &queue{}
	var choice string
	for true {
		fmt.Println("Enter your choice: ")
		fmt.Println("1-> Enque: ")
		fmt.Println("2-> Deque: ")
		fmt.Println("3-> size: ")
		fmt.Println("4-> Front: ")
		fmt.Println("5-> Display: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data int
			fmt.Println("Enter the element to queue: ")
			fmt.Scanln(&data)
			myQ.enque(data)
		case "2":
			myQ.deque()
		case "3":
			myQ.size()
		case "4":
			myQ.frontt()
		case "5":
			myQ.display()
		default:
			os.Exit(0)
		}
	}
}
