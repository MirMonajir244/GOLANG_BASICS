package main

import "fmt"

func roundInBubbleSort(arr []int, l int) int {
	var i, j int
	var swap_count int = 0
	for i = 0; i < l-1; i++ {
		for j = 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				swap_count++
			}
		}
	}
	return swap_count
}
func main() {
	var length int
	fmt.Println("Enter the array length: ")
	fmt.Scan(&length)
	var arr = make([]int, length+1)
	var k int
	for k = 0; k <= length; k++ {
		fmt.Scanln(&arr[k])
	}
	var count int = roundInBubbleSort(arr, length+1)
	fmt.Println("Total numbers of round needed to swap is: ", count)
}
