package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var result int
	slice := []rune{}

	for _, input := range s {
		for i, a := range slice {
			if a == input {
				slice = slice[i+1:]
				break
			}
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
