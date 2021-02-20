package main

import (
	"fmt"
)

/*
				bestSum

	>	Write a function `bestSum(targetSum, numbers)` that takes in a targetSum and an array of number as arguments.
	>	The function should return an array containing the shortest combination of numbers that add up to exactly the targetSum.
	>	If there is no combination that adds up to the targetSum, then return nil.
	>	If there is a tie for the shortest combination, you may return any one of the shortest.
	>	You may use an element of the array as many times as needed.
	>	You may assume that all input numbers are nonnegative.

*/

func bestSumMemo(targetSum int, numbers []int, memo map[int][]int) []int {
	if adds, ok := memo[targetSum]; ok {
		return adds
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	var shortestPath []int
	for _, num := range numbers {
		if adds := bestSumMemo(targetSum-num, numbers, memo); adds != nil {
			adds = append(adds, num)
			if shortestPath == nil {
				shortestPath = adds
			} else if len(adds) < len(shortestPath) {
				shortestPath = adds
			}
		}
	}
	memo[targetSum] = shortestPath
	return shortestPath
}

func evalBestSumMemo(targetSum int, numbers []int) string {
	if addends := bestSumMemo(targetSum, numbers, make(map[int][]int)); addends != nil {
		return fmt.Sprintf("%v", addends)
	}
	return "null"
}

func bestSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	var shortestPath []int
	for _, num := range numbers {
		if adds := bestSum(targetSum-num, numbers); adds != nil {
			adds = append(adds, num)
			if shortestPath == nil {
				shortestPath = adds
			} else if len(adds) < len(shortestPath) {
				shortestPath = adds
			}
		}
	}
	return shortestPath
}

func evalBestSum(targetSum int, numbers []int) string {
	if addends := bestSum(targetSum, numbers); addends != nil {
		return fmt.Sprintf("%v", addends)
	}
	return "null"
}

func main() {
	fmt.Println(evalBestSum(0, []int{2, 4}))              // []
	fmt.Println(evalBestSum(7, []int{}))                  // null
	fmt.Println(evalBestSum(7, []int{2, 4}))              // null
	fmt.Println(evalBestSum(7, []int{2, 3, 5, 7}))        // [7]
	fmt.Println(evalBestSum(8, []int{2, 3, 5}))           // [3 5]
	fmt.Println(evalBestSumMemo(300, []int{7, 14}))       // null
	fmt.Println(evalBestSumMemo(308, []int{7, 14}))       // [14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14 14]
	fmt.Println(evalBestSumMemo(100, []int{1, 2, 5, 25})) // [25 25 25 25]
}
