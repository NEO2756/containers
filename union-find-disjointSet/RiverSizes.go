package main

import (
	"fmt"
)

func main() {

	RiverSizes([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
		{1, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
	})
}

var parent []int

//var visited []bool
//var count int

func find(x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = find(parent[x])
	return parent[x]
}

func union(x, y int) {
	px := find(x)
	py := find(y)
	if px != py {
		parent[py] = px
	}
}
func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	r := len(matrix)
	c := len(matrix[0])

	parent = make([]int, r*c)
	freq := make([]int, r*c)

	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			fmt.Println(i, j)
			if i-1 >= 0 && matrix[i-1][j] == 1 {
				union(c*i+j, c*(i-1)+j)
			}
			if j-1 >= 0 && matrix[i][j-1] == 1 {
				union(c*i+j, (c*i + j - 1))
			}
		}
	}

	ans := make([]int, 0)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 1 {
				x := find(c*i + j)
				freq[x]++
			}
		}
	}
	for i := 0; i < r*c; i++ {
		if freq[i] > 0 {
			ans = append(ans, freq[i])
		}
	}
	fmt.Println(ans)
	return ans
}
