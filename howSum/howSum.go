package main

import (
	"fmt"
)

/*
				howSum

	>	Write a function `howSum(targetSum, numbers)` that takes in a targetSum and an array of number as arguments.
	>	The function should return an array containing any combination of elements that add up to exactly the targetSum.
	>	If there is no combination that adds up to the targetSum, then return nil.
	>	If there are multiple combinations possible, you may return any single one.
	>	You may use an element of the array as many times as needed.
	>	You may assume that all input numbers are nonnegative.

*/

func howSumMemo(targetSum int, numbers, addends []int, memo map[int][]int) []int {
	if adds, ok := memo[targetSum]; ok {
		return adds
	}
	if targetSum == 0 {
		return addends
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		if adds := howSumMemo(targetSum-num, numbers, append(addends, num), memo); adds != nil {
			memo[targetSum] = adds
			return adds
		}
	}

	memo[targetSum] = nil
	return nil
}

func evalHowSumMemo(targetSum int, numbers []int) string {
	if addends := howSumMemo(targetSum, numbers, []int{}, make(map[int][]int)); addends != nil {
		return fmt.Sprintf("%v", addends)
	}
	return "null"
}

func howSum(targetSum int, numbers, addends []int) []int {
	if targetSum == 0 {
		return addends
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		if addends := howSum(targetSum-num, numbers, append(addends, num)); addends != nil {
			return addends
		}
	}

	return nil
}

func evalHowSum(targetSum int, numbers []int) string {
	if addends := howSum(targetSum, numbers, []int{}); addends != nil {
		return fmt.Sprintf("%v", addends)
	}
	return "null"
}

func main() {
	fmt.Println(evalHowSumMemo(0, []int{2, 4}))       // []
	fmt.Println(evalHowSumMemo(7, []int{}))           // null
	fmt.Println(evalHowSumMemo(7, []int{2, 4}))       // null
	fmt.Println(evalHowSumMemo(7, []int{2, 3, 5, 7})) // :)
	fmt.Println(evalHowSumMemo(300, []int{7, 14}))    // null
	fmt.Println(evalHowSumMemo(308, []int{7, 14}))    // :)
}
