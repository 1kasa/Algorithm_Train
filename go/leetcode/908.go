package leetcode

import "slices"

func smallestRangeI(nums []int, k int) int {
    return max(0,slices.Max(nums)-slices.Min(nums)-2*k)
}