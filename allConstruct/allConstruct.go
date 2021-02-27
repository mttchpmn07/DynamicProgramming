package main

import (
	"fmt"
	"strings"
)

/*
				allConstruct

	>	Write a function `allConstruct(target, wordBank)` that accepts a target string and an array of strings.
	>	The function should return a 2D array containing all of the ways that the `target` can be constructed by concatenating
		elements of the `wordBank` array. Each element of the 2D array should represent one combination that constructs the `target`.
	>	You may reuse elements of `wordBank` as many times as needed.

*/

func prepend(a []string, b string) []string {
	return append([]string{b}, a...)
}

func allConstructMemo(target string, wordBank []string, memo map[string][][]string) [][]string {
	if v, ok := memo[target]; ok {
		return v
	}
	if len(target) == 0 {
		return nil
	}
	waysToConstruct := [][]string{}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			ways := allConstructMemo(strings.TrimPrefix(target, word), wordBank, memo)
			for i, way := range ways {
				ways[i] = prepend(way, word)
			}
			if ways == nil {
				ways = [][]string{{word}}
			}
			for _, way := range ways {
				waysToConstruct = append(waysToConstruct, way)
			}
		}
	}

	memo[target] = waysToConstruct
	return waysToConstruct
}

func allConstruct(target string, wordBank []string) [][]string {
	if len(target) == 0 {
		return nil
	}
	waysToConstruct := [][]string{}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			ways := allConstruct(strings.TrimPrefix(target, word), wordBank)
			for i, way := range ways {
				ways[i] = prepend(way, word)
			}
			if ways == nil {
				ways = [][]string{{word}}
			}
			for _, way := range ways {
				waysToConstruct = append(waysToConstruct, way)
			}
		}
	}

	return waysToConstruct
}

func main() {
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef"}))
	fmt.Println(allConstruct("", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(allConstruct("abcdef", []string{}))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(allConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeeee",
		"eeeeee",
		"eeeeeee",
		"eeeeeeee",
	}, make(map[string][][]string)))
}
