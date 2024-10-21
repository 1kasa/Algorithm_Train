package leetcode

import (
	"math"
	"slices"
)

func minimumAverage(nums []int) float64 {
	slices.Sort(nums)
	ans := math.MaxInt64
	for i,n:=0,len(nums);i<n/2;i++ {
		ans =min(ans,nums[i]+nums[n-i-1])
	}
	return float64(ans)/2
}