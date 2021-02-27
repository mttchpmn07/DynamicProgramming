package main

import (
	"fmt"
	"strings"
)

/*
				countConstruct

	>	Write a function `countConstruct(target, wordBank)` that accepts a target string and an array of strings.
	>	The function should return the number of ways that the `target` can be constructed by concatenating elements of the `wordBank` array.
	>	You may reuse elements of `wordBank` as many times as needed.

*/

func countConstructMemo(target string, wordBank []string, memo map[string]int) int {
	if i, ok := memo[target]; ok {
		return i
	}
	if len(target) == 0 {
		return 1
	}
	sum := int(0)
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			memo[target] = countConstructMemo(strings.TrimPrefix(target, word), wordBank, memo)
			sum += memo[target]
		}
	}
	memo[target] = sum
	return sum
}

func countConstruct(target string, wordBank []string) int {
	if len(target) == 0 {
		return 1
	}
	sum := int(0)
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			sum += countConstruct(strings.TrimPrefix(target, word), wordBank)
		}
	}
	return sum
}

func main() {
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef"}))
	fmt.Println(countConstruct("", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(countConstruct("abcdef", []string{}))
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(countConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeeee",
		"eeeeee",
		"eeeeeee",
		"eeeeeeee",
	}, make(map[string]int)))
}
