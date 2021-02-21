package main

import (
	"fmt"
	"strings"
)

/*
				canConstruct

	>	Write a function `canConstruct(target, wordBank)` that accepts a target string and an array of strings.
	>	The function should return a boolean indicating whether or not the `target` can be constructed by concatenating elements of the `wordBank` array.
	>	You may reuse elements of `wordBank` as many times as needed.

*/

func canConstructMemo(target string, wordBank []string, memo map[string]bool) bool {
	if b, ok := memo[target]; ok {
		return b
	}
	if len(target) == 0 {
		return true
	}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			if canConstructMemo(strings.TrimPrefix(target, word), wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}

func canConstruct(target string, wordBank []string) bool {
	if len(target) == 0 {
		return true
	}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			if canConstruct(strings.TrimPrefix(target, word), wordBank) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("abcdef", []string{}))
	fmt.Println(canConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeeee",
		"eeeeee",
		"eeeeeee",
		"eeeeeeee",
	}, make(map[string]bool)))
}
