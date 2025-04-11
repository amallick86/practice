package utils

import (
	"fmt"
	"sort"
)

func FindMaxAndMin(arr []int) (int, int) {
	fmt.Println("The given array is:", arr)
	sort.Ints(arr)
	max := arr[0]
	min := arr[len(arr)-1]
	println("The maximum number is:", max)
	println("The minimum number is:", min)
	return max, min
}
