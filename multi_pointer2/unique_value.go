package main

import "fmt"

func countUniqueValues(arr []int) int {
	start := 0;
	if len(arr)  < 1{
		return 0
	}
	for i := 1; i < len(arr); i++ {
		if arr[start] != arr[i] {
			start++;
			arr[start] = arr[i];
		} 
	}
	return start + 1;
}


func main() {
	fmt.Println(countUniqueValues([]int{1, 1, 1, 1, 1, 2}));
	fmt.Println(countUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}));
	fmt.Println(countUniqueValues([]int{-2, -1, -1, 0, 1}));
	fmt.Println(countUniqueValues([]int{}));
}