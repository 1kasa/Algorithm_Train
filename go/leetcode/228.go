package leetcode

import "strconv"

func summaryRanges(nums []int) (ans []string) {
	n := len(nums)
	for i := 0; i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}
