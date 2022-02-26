package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64

	// merge the arguments
	mergedArray := append(nums1, nums2...)

	// sorting by sort package
	sort.Slice(mergedArray, func(i, j int) bool { return mergedArray[i] < mergedArray[j] })

	if sliceLen := len(mergedArray); sliceLen%2 == 0 {
		// case of even
		result = float64(mergedArray[sliceLen/2-1]+mergedArray[sliceLen/2]) / 2.0
	} else {
		// case of odd
		result = float64(mergedArray[sliceLen/2])
	}
	return result
}

func main() {
	fmt.Printf("%v\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
