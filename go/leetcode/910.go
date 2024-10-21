package leetcode

import "slices"

func smallestRangeII(nums []int, k int) int {
	slices.Sort(nums)
	ans := nums[len(nums)-1] - nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		mx := max(nums[i-1]+k,nums[n-1]-k)
		mn := min(nums[0]+k,nums[i]-k)
		ans = min(ans,mx-mn)
	}
	return ans
}
