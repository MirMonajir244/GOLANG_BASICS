package main

import "fmt"

func findClosest(arr []int, n int, target int) int {
	if target <= arr[0] {
		return arr[0]
	}
	if target >= arr[n-1] {
		return arr[n-1]
	}

	var l, h, mid int = 0, n, 0
	for l < h {
		mid = (l + h) / 2
		if arr[mid] == target {
			return arr[mid]
		}
		//If target is less than array element, then search in left
		if target < arr[mid] {

			//If target is greater than previousto mid, return closest of two
			if mid > 0 && target > arr[mid-1] {
				return getClosest(arr[mid-1], arr[mid], target)
			}
			h = mid
		} else { // If target is greater than mid search right
			if mid < n-1 && target < arr[mid+1] {
				return getClosest(arr[mid], arr[mid+1], target)
				l = mid + 1
			}
		}
	}
	return arr[mid]
}

/*
 Method to compare which one is the more close.
 We find the closest by taking the difference
 between the target and both values
*/
func getClosest(val1, val2, val3 int) int {
	if val3-val1 >= val2-val3 {
		return val2
	} else {
		return val1
	}
}
func main() {
	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)

	fmt.Println("Enter array Elements: ")
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}

	var target int
	fmt.Println("Enter the Target Numer: ")
	fmt.Scanln(&target)

	fmt.Println("Closest number is: ")
	fmt.Println(findClosest(arr, size, target))
}

/*
Applied here Binary search Algorithm
Input:-
[2, 42, 82, 122, 162, 202, 242, 282, 322, 362]
target-80

Output:-


*/
