package leetcode

func maximumLength(nums []int, k int) int {
	// dp[v][j] - 表示以数字 v 结尾，并且有 j 次不同的“好”子序列的长度
	dp := make(map[int][]int, len(nums))
	// mx[j] - 表示在不同次数为 j 时的全局最长子序列
	mx := make([]int, k+2)

	for _, v := range nums {
		// 初始化 dp[v] 数组，长度为 k+1，表示不同次数为 0 到 k
		if dp[v] == nil {
			dp[v] = make([]int, k+1)
		}
		f := dp[v]

		// 从 k 到 0 更新 dp[v]，并同时更新全局最长 mx
		for j := k; j >= 0; j-- {
			// 更新 f[j]，即以 v 结尾并且有 j 次不同的最长“好”子序列
			f[j] = max(f[j], mx[j]) + 1
			// 更新 mx[j+1]，表示允许 j+1 次不同的全局最长长度
			mx[j+1] = max(mx[j+1], f[j])
		}
	}

	// 返回 mx[k+1]，即允许 k 次不同的“好”子序列的最长长度
	return mx[k+1]
}