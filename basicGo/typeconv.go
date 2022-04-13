package main

import (
	"fmt"
	"strconv"
)

//implicit type conversion in Go is not possible
// for converting any types of variable into string we have to use strconv package

func main() {
	var i int
	i = 20
	var j float32
	j = float32(i)
	fmt.Printf("%v %T\n", j, j)
	k := 40
	s := strconv.Itoa(k)
	fmt.Printf("%v %T\n", s, s)
	l, _ := strconv.Atoi("4")
	fmt.Printf("%v %T\n", l, l)
	q := strconv.Quote("Hello, হ্যালো")
	fmt.Printf("%v %T", q, q)
}
