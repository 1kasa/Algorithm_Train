package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	d1, d2 := 0, 0
	for i, v := range distance {
		if start <= i && i < destination {
			d1 += v
		} else {
			d2 += v
		}
	}
	return min(d1, d2)
} 
