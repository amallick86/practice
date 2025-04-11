package utils

import "fmt"

func CheckPalindrome(s string) bool {
	fmt.Println("The given string is:", s)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			fmt.Println("The string is not a palindrome")
			return false
		}
	}
	fmt.Println("The string is a palindrome")
	return true
}
