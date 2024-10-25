package leetcode

import "sort"

func mergeInterval(intervals [][]int) [][]int {
	sort.Slice(intervals,func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	ans = append(ans, intervals[0])
	for _, x:= range intervals[1:] {
		if ans[len(ans)-1][1] < x[0] {
			ans = append(ans, x)
		}else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], x[1])
		}
	}
	return ans
}