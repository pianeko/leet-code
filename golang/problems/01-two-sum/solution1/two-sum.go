package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// initialize
	var ans []int

loop:
	// loop based on nums
	for i, v := range nums {
		// loop based on sliced nums
		for j, w := range nums[i+1:] {
			if target == v+w {
				ans = []int{i, i + 1 + j}
				break loop
			}
		}
	}

	return ans
}

func main() {
	ans := twoSum([]int{3, 2, 4}, 6)
	fmt.Printf("%d\n", ans)
}
