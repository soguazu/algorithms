package main

import (
	"fmt"
	"math"
)

func same(a, b []int) bool {
	frequencyCounter1 := map[int]int{}
	frequencyCounter2 := map[int]int{}

	for _, value := range a {
		if _, ok := frequencyCounter1[value]; ok {
			frequencyCounter1[value]++
		} else {
			frequencyCounter1[value] = 1
		}
	}

	for _, value := range b {
		if _, ok := frequencyCounter2[value]; ok {
			frequencyCounter2[value]++
		} else {
			frequencyCounter2[value] = 1
		}
	}

	for key := range frequencyCounter1 {
		temp := int(math.Pow(float64(key), 2))
		
		if  _, ok := frequencyCounter2[temp]; !ok {
			return false
		}
		if frequencyCounter2[temp] != frequencyCounter1[key] {
			return false
		}
	}

	return true
}

func main() {
	arr1 := []int{1, 2, 3, 2, 5, 0}
	arr2 := []int{9, 1, 4, 4, 25, 0}
	fmt.Println(same(arr1, arr2))
}


