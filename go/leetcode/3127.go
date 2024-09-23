package leetcode

import ()

func canMakeSquare(grid [][]byte)bool {
	checkFunc := func(i,j int)bool {
		cnt:=[2]int{}
		cnt[grid[i][j]&1]++
		cnt[grid[i][j+1]&1]++
		cnt[grid[i+1][j]&1]++
		cnt[grid[i+1][j+1]&1]++
		return cnt[0]!=2
	}
	return checkFunc(0,0) || checkFunc(0,1) || checkFunc(1,0) || checkFunc(1,1)
}