package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"practice/utils"
)

func main() {
	problems := map[int]string{
		1:  "Find Non-Repeating Number",
		2:  "Reverse String",
		3:  "Check Palindrome",
		4:  "Find Maximum and Minimum",
		5:  "Binary Search",
		6:  "Merge Two Sorted Arrays",
		7:  "Find First and Last Position of Element in Sorted Array",
		8:  "Implement Queue using Stack",
		9:  "Detect Cycle in a Linked List",
		10: "Find kth Largest Element in Unsorted Array",
		11: "Valid Parentheses",
		12: "Implement LRU Cache",
		13: "Find Missing Number",
		14: "Rotate Array",
		15: "Count Islands",
		16: "Implement Trie (Prefix Tree)",
		17: "Longest Increasing Subsequence",
		18: "Coin Change Problem",
		19: "Word Break Problem",
		20: "Implement Heap Sort",
	}
	// Extracting the keys into a slice
	keys := make([]int, 0, len(problems))
	for k := range problems {
		keys = append(keys, k)
	}

	// Sorting the keys slice
	sort.Ints(keys)
	fmt.Println("==============================")
	fmt.Println("|  Golang DSA Practice Menu   |")
	fmt.Println("==============================")
	for _, key := range keys {
		fmt.Printf("%2d. %s\n", key, problems[key])
	}
	fmt.Println("------------------------------")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 0, 4, 5, 6, 7, 8}
		utils.NonRepeatingNumber(arr)
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a string to reverse: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		utils.ReverseString(s)
	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a string to check if it's a palindrome: ")
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		utils.CheckPalindrome(s)
	case 4:
		arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		utils.FindMaxAndMin(arr)
	default:
		fmt.Println("Invalid option.")
	}
}
