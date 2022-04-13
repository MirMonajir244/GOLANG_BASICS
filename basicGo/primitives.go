package main

import "fmt"

//Mainly 3 types of primitives we have in golang
// 1) Boolean type, 2)Numeric types-> int, float, complex and 3)text type
// in golang we can't perform arithmetic operations on differnet int types. ex- (int8+int32) impossible
//in golang by default the value of any not assigned variable is 0
func main() {
	//int
	a := 20
	b := 10
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	//important example
	var a1 int64 = 10
	var a2 int8 = 9
	fmt.Println(a1 + int64(a2))

	fmt.Println(64 >> 4) // 64 divide by 2 4 times
	fmt.Println(8 << 2)  // 8 multiplied by 2 2 times

	//complex

	var cm complex64 = 1 + 2i
	var cm1 complex64 = complex(2, 3)
	fmt.Println(cm)
	fmt.Println(cm1)
	fmt.Println(cm + cm1)

	//boolean

	var t bool = true
	fmt.Println(t)
	check := 1 == 1
	check1 := 1 == 2
	fmt.Println(check, check1)

	//text type
	// string-> UTF8 and rune->UTF32
	//strings are immutable
	str := "My Name is Mir Monajir"
	fmt.Println(str)
	fmt.Println(string(str[5]))
	str1 := "  Golang Developer"
	fmt.Println(str + str1)

}
