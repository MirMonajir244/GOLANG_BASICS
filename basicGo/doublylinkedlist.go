package main

import (
	"fmt"
	"os"
)

type node struct {
	data int
	prev *node
	next *node
}
type doublyLinkedlist struct {
	head *node
	tail *node
	len  int
}

func (dl *doublyLinkedlist) InsertEnd(val int) {
	var tmp = &node{}
	tmp.data = val
	tmp.prev = nil
	tmp.next = nil
	if dl.head == nil && dl.tail == nil {
		dl.head = tmp
		dl.tail = tmp
	} else {
		tmp.prev = dl.tail
		dl.tail.next = tmp
		dl.tail = tmp
	}
	dl.len++
}
func (dl *doublyLinkedlist) InsertBegin(val int) {
	var tmp = &node{}
	tmp.data = val
	tmp.prev = nil
	tmp.next = nil
	if dl.head == nil && dl.tail == nil {
		dl.head = tmp
		dl.tail = tmp
	} else {
		tmp.next = dl.head
		dl.head.prev = tmp
		dl.head = tmp
	}
	dl.len++
}
func (dl *doublyLinkedlist) display() {
	var l *node = dl.head
	for l != nil {
		fmt.Printf("%d ->", l.data)
		l = l.next
	}
}
func (dl *doublyLinkedlist) displayReverse() {
	var l *node = dl.tail
	for l != nil {
		fmt.Printf("%d ->", l.data)
		l = l.prev
	}
}
func (dl *doublyLinkedlist) deleteFromBegin() {
	if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head = dl.head.next
		dl.head.prev = nil
	}
	dl.len--
}
func (dl *doublyLinkedlist) deleteFromEnd() {
	if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
	}
	dl.len--
}
func (dl *doublyLinkedlist) isPresent(x int) bool {
	for dl.head != nil {
		if dl.head.data == x {
			return true
		}
		dl.head = dl.head.next
	}
	return false
}

func main() {
	var dlist = &doublyLinkedlist{}
	var choice string
	for true {
		fmt.Println("Enter your Choice: ")
		fmt.Println("1. Insert the data at END ")
		fmt.Println("2. Insert the data at Begining ")
		fmt.Println("3. Display/Traverse in head to tail ")
		fmt.Println("4. Display/Traverse in tail to head ")
		fmt.Println("5. Delete from the Begining ")
		fmt.Println("6. Delete from the End ")
		fmt.Println("7. Find a data from DoublyLinkedlist ")
		fmt.Println("8. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			fmt.Println("Enter the data: ")
			var d int
			fmt.Scanln(&d)
			dlist.InsertEnd(d)
		case "2":
			fmt.Println("Enter the data: ")
			var d2 int
			fmt.Scanln(&d2)
			dlist.InsertBegin(d2)
		case "3":
			dlist.display()
		case "4":
			dlist.displayReverse()
		case "5":
			dlist.deleteFromBegin()
		case "6":
			dlist.deleteFromEnd()
		case "7":
			var d3 int
			fmt.Println("Enter the data which you want to find: ")
			fmt.Scanln(&d3)
			res := dlist.isPresent(d3)
			if res == true {
				fmt.Println("Data is present")
			} else {
				fmt.Println("Data not present")
			}
		default:
			os.Exit(0)
		}
	}
}
