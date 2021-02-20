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

func howSumMemo(targetSum int, numbers []int, memo map[int][]int) []int {
	if adds, ok := memo[targetSum]; ok {
		return adds
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		if adds := howSumMemo(targetSum-num, numbers, memo); adds != nil {
			memo[targetSum] = append(adds, num)
			return memo[targetSum]
		}
	}

	memo[targetSum] = nil
	return nil
}

func evalHowSumMemo(targetSum int, numbers []int) string {
	if addends := howSumMemo(targetSum, numbers, make(map[int][]int)); addends != nil {
		return fmt.Sprintf("%v", addends)
	}
	return "null"
}

func howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		if adds := howSum(targetSum-num, numbers); adds != nil {
			return append(adds, num)
		}
	}

	return nil
}

func evalHowSum(targetSum int, numbers []int) string {
	if addends := howSum(targetSum, numbers); addends != nil {
		return fmt.Sprintf("%v", addends)
	}
	return "null"
}

func main() {
	fmt.Println(evalHowSumMemo(0, []int{2, 4}))       // []
	fmt.Println(evalHowSumMemo(7, []int{}))           // null
	fmt.Println(evalHowSumMemo(7, []int{2, 4}))       // null
	fmt.Println(evalHowSumMemo(7, []int{2, 3, 5, 7})) // [3 2 2]
	fmt.Println(evalHowSumMemo(8, []int{2, 3, 5}))    // [2, 2, 2, 2]
	fmt.Println(evalHowSumMemo(300, []int{7, 14}))    // null
	fmt.Println(evalHowSumMemo(308, []int{7, 14}))    // [7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7 7]
}
