package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head *node
	len  int
}

func (l *linkedList) Insert(val int) {
	n := node{data: val}
	n.data = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

/*
I used here slow and fast pointer approach
here slow=head and fast=head.next
*/
func isPalindrome(h *node) bool {
	if h == nil || h.next == nil {
		return true
	}
	slow := h
	fast := h.next
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	curr := slow.next
	slow.next = nil
	mid := reverseLinkedList(curr)
	for mid != nil && h != nil {
		if mid.data != h.data {
			return false
		}
		mid = mid.next
		h = h.next
	}
	return true
}
func reverseLinkedList(h *node) *node {
	if h == nil && h.next == nil {
		return h
	}
	p := h        // pointer to traverse the linked list
	var rst *node // pointer to save the final result linked list
	var q *node   // temporary pointer to save the next node

	for p != nil {
		q = p.next
		p.next = rst
		rst = p
		p = q
	}
	return rst
}
func main() {
	list := &linkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(1)
	res := isPalindrome(list.head)
	if res == true {
		fmt.Println("Linkedlist is paindrome")
	} else {
		fmt.Println("Linkedlist is not palindrome")
	}

}
