package main

import "fmt"

type node struct {
	data int
	next *node
}

func detectCountLoop(head *node) int {

	small := head
	large := head
	cnt := 0
	flag := 0

	for large != nil {
		small = (*small).next
		if large.next == nil {
			large = large.next
		} else {
			large = (*large).next.next
		}

		if large != nil {

			if small == large {
				flag = 1
				break
			}
		}
	}

	if flag == 1 {
		cnt++
		large = (*large).next

		for large != small {
			cnt++
			large = (*large).next
		}
	}

	return cnt
}

func find(head *node, element int) bool {
	temp1 := head

	for temp1 != nil {
		if (*temp1).data == element {
			return true
		}

		temp1 = (*temp1).next
	}

	return false
}

func main() {

	var head *node = nil
	var tail *node = nil
	for true {
		var data int
		var con int
		fmt.Println("Enter the data")
		fmt.Scanf("%d\n", &data)

		if head == nil {
			var temp node
			temp.data = data
			temp.next = nil
			head = &temp
			tail = &temp

		} else if find(head, data) {

			temp1 := head

			for temp1.data != data {
				temp1 = (*temp1).next
			}

			(*tail).next = temp1
			break

		} else {

			var temp node
			temp.data = data
			temp.next = nil

			var temp1 *node
			temp1 = head

			for (*temp1).next != nil {
				temp1 = (*temp1).next
			}

			(*temp1).next = &temp
			tail = &temp

		}

		fmt.Println("Do you want to continue [1/0]")
		fmt.Scanf("%d\n", &con)

		if con == 0 {
			break
		}

	}

	cnt := detectCountLoop(head)

	if cnt == 0 {
		fmt.Println("NO", -1)
	} else {
		fmt.Println("YES", cnt)
	}

}
