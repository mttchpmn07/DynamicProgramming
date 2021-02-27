package main

import "fmt"

/*
				allConstruct

	>	Write a function `allConstruct(target, wordBank)` that accepts a target string and an array of strings.
	>	The function should return a 2D array containing all of the ways that the `target` can be constructed by concatenating
		elements of the `wordBank` array. Each element of the 2D array should represent one combination that constructs the `target`.
	>	You may reuse elements of `wordBank` as many times as needed.

*/

func allConstruct(target string, wordBank []string) [][]string {
	/*
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
	*/

	return [][]string{}
}

func main() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef"}))
	fmt.Println(allConstruct("", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(allConstruct("abcdef", []string{}))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(allConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeeee",
		"eeeeee",
		"eeeeeee",
		"eeeeeeee",
	})) //, make(map[string]int)))

}
