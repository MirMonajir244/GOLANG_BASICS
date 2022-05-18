package main

import "fmt"

//Insertion sort is basic sorting algorithm it is good for partially sorted Arrays
func insertionSort(arr []int) {
	var i, j, key int
	n := len(arr)
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1
		for (j >= 0) && key < arr[j] {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
func printaRRY(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	var arr = []int{4, 5, 1, 3, 2}
	insertionSort(arr)
	fmt.Println("After sotring array is: ")
	printaRRY(arr)
}
