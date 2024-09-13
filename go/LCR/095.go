package lcr

func longestCommonSubsequence(text1 string,test2 string)int {
	m,n := len(text1),len(test2)
	dp := make([][]int,m+1)
	for i:=0;i<=m;i++ {
		dp[i] = make([]int,n+1)
	}
	for i,v1 := range text1 {
		for j,v2 := range test2 {
			if v1 == v2 {
				dp[i+1][j+1]=dp[i][j]+1
			}else {
				dp[i+1][j+1]=max(dp[i][j+1],dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}