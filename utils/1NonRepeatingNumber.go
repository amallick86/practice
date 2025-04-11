package utils

import "fmt"

func NonRepeatingNumber(arr []int) int {
	fmt.Println("The given array is:", arr)
	count := make(map[int]int)

	for _, v := range arr {
		count[v]++
	}
	for _, v := range arr {
		if count[v] == 1 {
			fmt.Println("The non repeating number is:", v)
			return v
		}
	}
	return -1
}
