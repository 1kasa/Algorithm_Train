package leetcode

func maximumSubsequenceCount(text string, pattern string) int64 {
	var ans int64
	x, y := pattern[0], pattern[1]
	cntX, cntY := 0, 0
	for i := range text {
		if text[i] == y {
            ans += int64(cntX)
			cntY++
		}
		if text[i]==x {
			cntX++
		}
	}
	return ans + int64(max(cntX,cntY))
}
