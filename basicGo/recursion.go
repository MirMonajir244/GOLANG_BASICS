package main

import (
	"fmt"
	"math"
)

func square(a float64) float64 {
	var res float64
	res = math.Pow(a, 2)
	return res
}
func qube(sqr func(float64) float64) float64 {
	var res1 float64
	res1 = math.Pow(sqr(2), 3)
	return res1
}
func main() {
	var a float64
	fmt.Scan(&a)
	fmt.Println(qube(square))
}
