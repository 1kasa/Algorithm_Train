package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	startX, startY := 0, 0	
	mid := n / 2
	count, offset := 1, 1
	for loop := n / 2; loop > 0; loop-- {
		i, j := startX, startY

		for j = startY; j < n-offset; j++ {
			res[startX][j] = count
			count++
		}

		for i = startX; i < n-offset; i++ {
			res[i][j] = count
			count++
		}

		for ; j > startY; j-- {
			res[i][j] = count
			count++
		}

		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}

		startX++
		startY++
		offset++
	}

	if n%2 == 1 {
		res[mid][mid] = count
	}
	return res
}
