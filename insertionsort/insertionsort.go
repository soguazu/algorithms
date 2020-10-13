package main

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	currentValue := 0
	index := 0
	for i := 1; i < len(arr); i++ {
		currentValue = arr[i]
		index = i
		for j := i - 1; j >= 0 && arr[j] > currentValue; j-- {
			arr[j + 1] = arr[j]
			index--
		}
		arr[index] = currentValue
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{4, 9, 3, 14, 11, 50, 34, 23, 12}))
}