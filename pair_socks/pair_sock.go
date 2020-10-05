package main

import (
	"fmt"
)


func socksMerchant(arr []int) int {
	temp := map[int]int{}
	counter := 0

	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; ok {
			temp[arr[i]]++
		} else {
			temp[arr[i]] = 1
		}
	}

	for _, value := range temp {
		if value % 2 != 0 && value > 1 {
			counter += (value - (value % 2)) / 2
		}

		if value % 2 == 0 {
			counter += value / 2
		}
	}
	return counter
	
}


func main() {
	arr := []int{10, 20, 20, 30, 10, 10, 10, 30, 50, 1020}
	fmt.Println(socksMerchant(arr))

}

