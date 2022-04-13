package main

import "fmt"

// WE can declare a variable using 3 types
// If we declare variable using captial letter means that variable is globally accessiable
var I int = 20

func main() {
	var a int
	a = 10
	fmt.Println(a)
	var b float64 = 6.7
	fmt.Println(b)
	c := "abcd"
	fmt.Println(c)
}
