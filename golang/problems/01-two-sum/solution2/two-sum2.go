package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// initialize
	var ans []int

loop:
	// loop based on nums
	for i := 0; i < len(nums)-1; i++ {
		// loop based on sliced nums
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				ans = []int{i, j}
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
