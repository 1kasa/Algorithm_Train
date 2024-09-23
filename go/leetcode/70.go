package leetcode

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
	dp := make([]int, n+1)
	dp[1]=1
	dp[0]=0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairsOther1(n int) int {
	if n == 1 {
		return 1
	}
	p, q, x := 1, 1, 0
	for i := 2; i <= n; i++ {
		x = p + q
		p = q
		q = x
	}
	return x
}

func climbStairsOther2(n int) int {
    p, q, r := 0, 0, 1
    for i := 1; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}