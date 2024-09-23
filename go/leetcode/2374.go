package leetcode

func edgeScore(edges []int) (ans int) {
	n := len(edges)
	score := make([]int, n)
	for i, v := range edges {
		score[v] += i
		if score[v] > score[ans] || (score[ans] == score[v] && v < ans) {
			ans = v
		}
	}
	return
}
