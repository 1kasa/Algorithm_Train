package leetcode

func combinationSum(candidates []int, target int) (ans [][]int) {
	res := make([]int, 0)
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), res...))
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			res = append(res, candidates[index])
			dfs(target-candidates[index], index)
			res = res[:len(res)-1]
		}
	}
	dfs(target, 0)
	return ans
}
