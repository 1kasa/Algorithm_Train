package leetcode

func findJudge(n int, trust [][]int) int {
	trustmp := make([]int, n+1)
	untrustmp := make([]int, n+1)
	for _, v := range trust {
		trustmp[v[1]]++
		untrustmp[v[0]]++
	}
	for i := 1; i <= n; i++ {
		if trustmp[i] == n-1 && untrustmp[i] == 0 {
			return i
		}
	}
	return -1
}
