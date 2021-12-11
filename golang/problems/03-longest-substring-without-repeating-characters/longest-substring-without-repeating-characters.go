package main

import (
	"fmt"
	"strings"
)

func contains(slice []string, s string) (int, bool) {
	for i, a := range slice {
		if a == s {
			return i + 1, true
		}
	}
	return 0, false
}

func lengthOfLongestSubstring(s string) int {
	var result int
	inputs := strings.Split(s, "")
	slice := []string{}

	for _, input := range inputs {
		index, isContain := contains(slice, input)
		if isContain {
			if len(slice) > result {
				result = len(slice)
			}
			slice = slice[index:]
		}
		slice = append(slice, input)
		if len(slice) > result {
			result = len(slice)
		}
	}
	return result
}

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf"))
}
