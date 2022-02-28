/**
 * Definition of Interval:

 */
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */

type intheap []int

func (h intheap) Len() int {
	return len(h)
}
func (h intheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h intheap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *intheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *intheap) Top() int {
	return (*h)[0]
}

func minMeetingRooms(intervals []*Interval) int {
	// Write your code here

	if len(intervals) == 0 {
		return 0
	}

	h := make(intheap, 0)

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	///fmt.Println(intervals[0], intervals[1])

	ans := 1
	heap.Push(&h, intervals[0].End)

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < h.Top() {
			ans++
		} else {
			heap.Pop(&h)
		}
		heap.Push(&h, intervals[i].End)
	}
	return ans
}

func main() {
	arr := []*Interval{
		{1, 5},
		{6, 10},
		{6, 11},
		{10, 15},
		{10, 16},
	}
	fmt.Println(minMeetingRooms(arr))
}
