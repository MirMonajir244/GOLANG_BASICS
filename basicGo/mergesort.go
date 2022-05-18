package main

import "fmt"

//Time complexity of merge sort is O(nlogn)

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2
	left_part := mergeSort(array[:mid])
	right_part := mergeSort(array[mid+1:])
	return merge(left_part, right_part)
}

func merge(array1 []int, array2 []int) []int {
	final_sorted := []int{}
	i := 0
	j := 0
	for i < len(array1) && j < len(array2) {
		if array1[i] < array2[j] {
			final_sorted = append(final_sorted, array1[i])
			i++
		} else {
			final_sorted = append(final_sorted, array2[j])
		}
	}
	for ; i < len(array1); i++ {
		final_sorted = append(final_sorted, array1[i])
	}
	for ; j < len(array2); j++ {
		final_sorted = append(final_sorted, array2[j])
	}
	return final_sorted
}
func printArray(A []int, sz int) {
	for i := 0; i < sz; i++ {
		fmt.Println(A[i])
	}
}
func main() {

	unsorted := []int{2, 5, 1, 3, 9, 10, 56, 11}
	sz := len(unsorted)
	fmt.Println("Given Usorted array is: ")
	printArray(unsorted, sz)
	mergeSort(unsorted)

	fmt.Println("After merge sort array is: ")
	printArray(unsorted, sz)
}
