package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var result int
	slice := []rune{}

	//loop based on inputs
	for _, input := range s {
		//check if the slice contains a value
		for i, a := range slice {
			if a == input {
				//extract the values after the corresponding value
				slice = slice[i+1:]
				break
			}
		}
		//add value
		slice = append(slice, input)
		//if the length of the slice is greater than the result, set the length
		if len(slice) > result {
			result = len(slice)
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf"))
}
