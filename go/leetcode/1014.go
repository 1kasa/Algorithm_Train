package leetcode

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	ans, mx := 0, values[0]+0
	for i := 1; i < n; i++ {
		ans = max(ans, mx+values[i]-i)
		mx = max(mx, values[i]+i)
	}
	return ans
}