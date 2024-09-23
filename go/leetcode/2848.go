package leetcode

import ()

func numberOfPoints(nums [][]int) int {
	C := 0
	for _, inteval := range nums {
		if inteval[1] > C {
			C = inteval[1]
		}
	}

	count := make([]int, C+1)
	for _, interval := range nums {
		for i := interval[0]; i < interval[1]; i++ {
			count[i]++
		}
	}

	ans := 0
	for i := 1; i <= C; i++ {
		if count[i] > 0 {
			ans++
		}
	}
	return ans
}

func numberOfPointsOther(nums [][]int) int {
	ans := 0
	maxLen := 0
	for _, v := range nums {
		maxLen = max(maxLen, v[1])
	}

	diff := make([]int, maxLen+2)
	for _, v := range nums {
		diff[v[0]]++
		diff[v[1]+1]--
	}

	s := 0
	for _, v := range diff {
		s += v
		if s > 0 {
			ans++
		}
	}
	return ans
}
