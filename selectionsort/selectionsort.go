package main

import (
	"fmt"
)

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		lowest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowest] {
				lowest = j
			}
		}

		// This condition is checked because if i  and lowest is the same,
		// They carry the same index, so there is no need to swap
		if i != lowest {
			arr[i], arr[lowest] = arr[lowest], arr[i]
		}
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{4, 9, 3, 14, 11, 50, 34, 23, 12}))
	fmt.Println(selectionSort([]int{1, 2, 3, 4, 5}))
}