package leetcode

func countCompleteDayPairs(hours []int) int {
	ans := 0
	for i := 1; i < len(hours); i++ {
		for j := 0; j < i; j++ {
			if (hours[i]+hours[j])%24 == 0 {
				ans++
			}
		}
	}
	return ans
}