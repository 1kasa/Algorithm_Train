package leetcode

func maxStrength(nums []int) int64 {
	mx, mn := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn = max(mx, max((mn*nums[i]), max(nums[i], mx*nums[i]))), min(mn, min(mx*nums[i], min(nums[i], mn*nums[i])))
	}
	return int64(mx)
}
