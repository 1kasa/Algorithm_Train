package leetcode

import (
	"container/heap"
	"sort"
)

// An Interval is an interval with a start and end time.
type Interval struct {
	Start int
	End   int
}

// A MinHeap is a min-heap of integers.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] } // MinHeap
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// minMeetingRooms calculates the minimum number of meeting rooms required.
func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	// Create a min-heap to track the end times of meetings
	h := &MinHeap{}
	heap.Init(h)

	// Sort the intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// Add the first meeting's end time to the heap
	heap.Push(h, intervals[0].End)

	// Iterate through the rest of the meetings
	for i := 1; i < len(intervals); i++ {
		// If the meeting starts after the earliest ending meeting in the heap
		if intervals[i].Start >= (*h)[0] {
			heap.Pop(h) // Remove the meeting that ended
		}

		// Add the current meeting's end time to the heap
		heap.Push(h, intervals[i].End)
	}

	// The size of the heap is the number of rooms required
	return h.Len()
}
