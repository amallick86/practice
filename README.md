# Golang Data Structures and Algorithms Practice

This repository contains various Data Structures and Algorithms problems implemented in Go.

## Problems

### 1. Find Non-Repeating Number
Given an array where every element appears twice except for one element, find that single element.
```go
// Example: In []int{2, 3, 4, 2, 3}, the non-repeating number is 4
```

### 2. Reverse String
Write a function that reverses a string.
```go
// Example: reverseString("hello") should return "olleh"
```
### 3. Check Palindrome
Write a function that checks if a given string is a palindrome.
```go
// Example: isPalindrome("racecar") should return true
```
### 4. Find Maximum and Minimum
Write a function that finds the maximum and minimum values in an array.
```go
// Example: findMaxMin([]int{1, 2, 3, 4, 5}) should return 5 and 1
```
### 5. Binary Search
Implement a binary search algorithm.
```go
// Example: binarySearch([]int{1, 2, 3, 4, 5}, 3) should return 2
``` 
### 6. Merge Two Sorted Arrays
Write a function that merges two sorted arrays into a single sorted array.
```go
// Example: mergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}) should return []int{1, 2, 3, 4, 5, 6}
```
### 7. Find First and last Position of Element in Sorted Array
Write a function that finds the first and last position of a given element in a sorted array.
```go
// Example: searchRange([]int{1, 2, 2, 2, 3, 4, 4}, 2) should return []int{1, 3}
```
### 8. Implement Queue using Stack
Implement a queue using two stacks.
```go
// Example: enqueue(1), enqueue(2), dequeue() should return 1
```
###9. Detect Cycle in a Linked List
Write a function that detects if a linked list contains a cycle.
```go
// Example: hasCycle(1 -> 2 -> 3 -> 4 -> 2) should return true
```
### 10. Find kth Largest Element in unsorted array
Write a function that finds the kth largest element in an unsorted array.
```go
// Example: findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) should return 5
```
### 11. Valid Parentheses
Write a function that checks if a given string of parentheses is valid.
```go
// Example: isValid("()") should return true
```
### 12. Impplement LRU Cache
Implement a Least Recently Used (LRU) cache.
```go
// Example: get(1), put(2, 2), get(1), put(3, 3), get(2) should return 1, nil, -1
```
### 13. Find Missing Number
Write a function that finds the missing number in an array of integers.
```go
// Example: findMissingNumber([]int{1, 2, 3, 5}) should return 4
```
### 14. Rotate Array
Write a function that rotates an array of integers to the right by a given number of steps.
```go
// Example: Rotating []int{1, 2, 3, 4, 5, 6, 7} by 3 steps gives []int{5, 6, 7, 1, 2, 3, 4}
```
### 15. Count Islands
Write a function that counts the number of islands in a given 2D grid.
```go
// Example: countIslands([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}) should return 3
// Example: In a grid with 1s representing land and 0s representing water, count connected 1s as islands
```
### 16. Implement Trie (Prefix Tree)
Implement a Trie (Prefix Tree) data structure.
```go
// Example: insert("apple"), search("apple"), startsWith("app") should return true, true, true
```
### 17. Longest Increasing Sunsequence
Write a function that finds the length of the longest increasing subsequence in an array.
```go
// Example: In []int{10, 9, 2, 5, 3, 7, 101, 18}, the longest increasing subsequence is [2, 3, 7, 101], with length 4
```
### 18. Coin Change Problem
Write a function that finds the minimum number of coins needed to make a given amount of change.
```go
// Example: coinChange([]int{1, 2, 5}, 11) should return 3 (11 = 5 + 5 + 1)
```
### 19. Word Break problem
Write a function that determines if a given string can be segmented into a space-separated sequence of one or more dictionary words.
```go
// Example: wordBreak("leetcode", []string{"leet", "code"}) should return true
```
### 20. Implement Heap Sort
Implement the Heap Sort algorithm.
```go
// Example: heapSort([]int{1, 4, 2, 3, 5, 2}) should return []int{1, 2, 2, 3, 4, 5}
```