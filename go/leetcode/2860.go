package leetcode

import (
	"slices"
)

func countWays(nums []int) int {
    slices.Sort(nums)
	n := len(nums)
	count := 0

	if nums[0] > 0 {
		count++
	}

	for x := 1; x < n; x++ {
		if nums[x-1] < x && x < nums[x] {
			count++
		}
	}

	if nums[n-1] < n {
		count++
	}
	return count
}
