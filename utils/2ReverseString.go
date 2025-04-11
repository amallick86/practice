package utils

import "fmt"

func ReverseString(s string) string {
	fmt.Println("The given string is:", s)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println("The reversed string is:", string(r))
	return string(r)
}
