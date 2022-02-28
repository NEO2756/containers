// https://leetcode.com/problems/largest-component-size-by-common-factor/
package main

import "fmt"

var parent []int
var rank []int
var ans int

func find(x int) int {

	if x == parent[x] {
		return x
	}

	return find(parent[x])
}

func union(x, y int) {

	px := find(x)
	py := find(y)

	if px != py {
		if rank[px] > rank[py] {
			parent[py] = px
			rank[px] += rank[py]
			ans = max(ans, rank[px])
		} else {
			parent[px] = py
			rank[py] += rank[px]
			ans = max(ans, rank[py])
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func largestComponentSize(nums []int) int {

	parent = make([]int, 20001)
	rank = make([]int, 20001)
	ans = 0

	for i := 0; i < 20001; i++ {
		parent[i] = i
		rank[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			m := min(nums[i], nums[j])
			for k := 2; k <= m; k++ {
				if nums[i]%k == 0 && nums[j]%k == 0 {
					union(nums[i], nums[j])
				}
			}
		}
	}

	fmt.Println(parent[:40])
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func main() {
	a := largestComponentSize([]int{99, 100, 69, 39, 14, 56, 91, 60})
	fmt.Println(a)
}
