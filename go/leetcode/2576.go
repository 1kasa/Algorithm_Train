package leetcode

import (
	"slices"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, n/2
	check := func(m int) bool {
		for i := 0; i < m; i++ {
			if nums[i]*2 > nums[n-m+i] {
				return false
			}
		}
		return true
	}
	for l < r {
		m := (l + r + 1) >> 1
		if check(m) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l * 2
}

func maxNumOfMarkedIndicesOther(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	pairs := sort.Search(n/2, func(k int) bool {
		k++
		for i, x := range nums[:k] {
			if 2*x > nums[n-k+i] {
				return true
			}
		}
		return false
	})
	return pairs * 2
}
