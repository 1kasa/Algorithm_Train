package leetcode

func countCompleteDayPairsII(hours []int) int64 {
	const H = 24
	ans := int64(0)
	cnt := [H]int{}
	for _, v := range hours {
		ans += int64(cnt[(H-v%H)%H])
		cnt[v%H]++
	}
	return ans
}
