package main

import (
	"fmt"
)

type point struct {
	val, frq int
}

var heap []point
var size int

func parent(p int) int {
	return (p - 1) / 2
}

func left(p int) int {
	return 2*p + 1
}
func right(p int) int {
	return 2*p + 2
}

func insert(p point) {

	heap = append(heap, p)
	size = len(heap)

	pos := size - 1
	for pos > 0 && heap[pos].frq > heap[parent(pos)].frq {
		heap[pos].frq, heap[parent(pos)].frq = heap[parent(pos)].frq, heap[pos].frq
		heap[pos].val, heap[parent(pos)].val = heap[parent(pos)].val, heap[pos].val
		pos = parent(pos)
	}
}

func heapify(pos int) {

	l := left(pos)
	r := right(pos)

	if l > size && r > size {
		return
	}

	max := l
	if l <= size {
		max = l
	}
	if r <= size {
		if heap[max].frq < heap[r].frq {
			max = r
		}
	}

	if heap[max].frq > heap[pos].frq {
		heap[max].frq, heap[pos].frq = heap[pos].frq, heap[max].frq
		heap[max].val, heap[pos].val = heap[pos].val, heap[max].val
		pos = max
		heapify(pos)
	}

}

func delete() point {
	p := heap[0]
	heap[0] = heap[size-1]
	size--

	heapify(0)
	return p
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	ans := make([]int, 0)
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		insert(point{val: k, frq: v})
	}

	fmt.Println(heap)
	for i := 0; i < k; i++ {
		ans = append(ans, delete().val)
	}

	fmt.Println(ans)
	return ans
}
