package main

import "fmt"

type node struct {
	data int
	next *node
}
type singlyLinkedList struct {
	len  int
	head *node
}

func initList() *singlyLinkedList {
	return &singlyLinkedList{}
}
func (s *singlyLinkedList) AddFront(data int) {
	node := &node{
		data: data,
	}
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.len++
	return
}

func (s *singlyLinkedList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

//Function to convert singly linked list to circular linked list
func (s *singlyLinkedList) ToCircular() {
	current := s.head
	for current.next != nil {
		current = current.next
	}
	current.next = s.head
}

func (s *singlyLinkedList) IsCircular() bool {
	if s.head == nil {
		return true
	}
	current := s.head.next
	for current.next != nil && current != s.head {
		current = current.next
	}
	return current == s.head
}

func main() {
	List := initList()
	fmt.Println("Add data: ")
	List.AddFront(2)
	List.AddFront(3)
	List.AddFront(4)
	List.AddFront(2)
	List.AddFront(1)
	err := List.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}
	List.ToCircular()
	isCircular := List.IsCircular()
	if isCircular == true {
		fmt.Println("YES")
	}
	fmt.Println("Size: ", List.len)
}

/*
A linked list is circular if all the nodes are connected in the form of a cylce, means unlike a singly linkedlist last node pointing to the null circular will point to the head node
steps=>
first I have create a linked list.
After that, I have convert the linked list into a circular linked list.
Then check if it is circular It prints YES with length

Input:-1->2->3->4->2
output:-
PS C:\Users\RD\Desktop\GolangCORIOLIS\basicGo> go run circularLinkedList.go
Add data:
1
2
4
3
2
YES
Size:  5
*/
