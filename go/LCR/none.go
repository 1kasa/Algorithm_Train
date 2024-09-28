package lcr

func longestPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if dp[r][c] != 0 {
			return dp[r][c]
		}
		maxLength := 1
		directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 下，上，右，左
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && abs(matrix[nr][nc]-matrix[r][c]) == 1 {
				maxLength = max(maxLength, 1+dfs(nr, nc))
			}
		}
		dp[r][c] = maxLength
		return maxLength
	}

	maxPath := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			maxPath = max(maxPath, dfs(r, c))
		}
	}

	return maxPath
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
