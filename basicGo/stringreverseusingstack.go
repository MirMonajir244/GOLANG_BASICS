package main

import (
	"bufio"
	"fmt"
	"os"
)

type stringStack struct {
	my_string []byte
}

func (s *stringStack) push(data byte) {
	s.my_string = append(s.my_string, data)
}

func (s *stringStack) top() byte {
	return s.my_string[len(s.my_string)-1]
}

func (s *stringStack) pop() {

	if len(s.my_string) != 0 {
		s.my_string = s.my_string[0 : len(s.my_string)-1]
	} else {
		fmt.Println("Stack is Empty")
	}
}

func main() {

	var s *stringStack = new(stringStack)
	var str string
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string: ")
	sc.Scan()
	str = sc.Text()
	for i := 0; i < len(str); i++ {
		s.push(str[i])
	}
	for len(s.my_string) != 0 {
		fmt.Printf("%s", string(s.top()))
		s.pop()
	}

}

/*
Input:- "reverse me"
Output:-
PS C:\Users\RD\Desktop\GolangCORIOLIS\basicGo> go run stringreverseusingstack.go
Enter the string:
reverse me
em esrever
*/
