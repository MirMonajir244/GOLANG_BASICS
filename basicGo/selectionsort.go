package main

import "fmt"

func selectionSort(arr []int) {
	var i, j, min_index int
	n := len(arr)
	for i = 0; i < n-1; i++ {
		min_index = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min_index]
		arr[min_index] = temp
	}
}
func printarray(arr []int) {
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	var arr = []int{4, 5, 1, 3, 2}
	selectionSort(arr)
	fmt.Println("After sorting array is: ")
	printarray(arr)
}
