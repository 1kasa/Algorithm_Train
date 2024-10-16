package leetcode

import (
	"slices"
	"sort"
)

func minimumTime(time []int, totalTrips int) int64 {
	minT, maxT := slices.Min(time), slices.Max(time)
	avg := (totalTrips-1)/len(time) + 1
	left := minT * avg
	right := min(minT*totalTrips, maxT*avg)
	return int64(left + sort.Search(right-left, func(x int) bool {
		sum := 0
		x += left
		for _, t := range time {
			sum += x / t
			if sum >= totalTrips {
				return true
			}
		}
		return false
	}))
}
