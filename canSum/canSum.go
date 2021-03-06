package main

import "fmt"

/*
				canSum

	>	Write a function `canSum(targetSum, numbers)` that takes in a targetSum and an array of number as arguments.
	>	The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers from the array.
	>	You may use an element of the array as many times as needed.
	>	You may assume that all input numbers are nonnegative.

*/

func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, num := range numbers {
		if canSum(targetSum-num, numbers) {
			return true
		}
	}

	return false
}

func canSumMem(targetSum int, numbers []int, memo map[int]bool) bool {
	if b, ok := memo[targetSum]; ok {
		return b
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {
		if canSumMem(targetSum-num, numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false
}

func main() {
	fmt.Println(canSumMem(7, []int{2, 3}, make(map[int]bool)))       // true
	fmt.Println(canSumMem(7, []int{5, 3, 4, 7}, make(map[int]bool))) // true
	fmt.Println(canSumMem(7, []int{2, 4}, make(map[int]bool)))       // false
	fmt.Println(canSumMem(8, []int{2, 3, 5}, make(map[int]bool)))    // true
	fmt.Println(canSumMem(300, []int{7, 14}, make(map[int]bool)))    // false
}
