package leetcode

import ()

func differenceOfSum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
		for v > 0 {
			sum -= v % 10
			v /= 10
		}
	}
	return
}