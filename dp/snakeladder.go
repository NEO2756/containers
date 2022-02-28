package main

import (
	"fmt"
)

var saanp map[int]int
var seedi map[int]int

/*
moves[2] = 21
moves[4] = 7
moves[10] = 25
moves[19] = 28

# Snakes
moves[26] = 0
moves[20] = 8
moves[16] = 3
moves[18] = 6
*/

func fillSaanp() {
	saanp = make(map[int]int)
	seedi = make(map[int]int)
	saanp[26] = 0
	saanp[20] = 8
	saanp[16] = 3
	saanp[18] = 6

	//seedi[23] = 2
	seedi[28] = 13
	//seedi[25] = 10
	seedi[28] = 19

}

func saanpseedi() int {
	fillSaanp()
	// dp[i] == mini moves to reach ith location
	dp := make([]int, 31)

	dp[0] = 999999
	dp[1], dp[2], dp[3], dp[4], dp[5], dp[6] = 1, 1, 1, 1, 1, 1

	for i := 7; i < 31; i++ {

		if _, ok := saanp[i]; ok {
			dp[i] = 999999
			continue
		}

		if v, ok := seedi[i]; ok {
			tmp := min(min(dp[i-1], dp[i-2]), min(min(dp[i-3], dp[i-4]), min(dp[i-5], dp[i-6])))
			if tmp < dp[v] {
				fmt.Println("Corenr", v)
				dp[i] = 1 + dp[tmp]
			} else {
				dp[i] = dp[v]
			}
		} else {
			dp[i] = 1 + min(min(dp[i-1], dp[i-2]), min(min(dp[i-3], dp[i-4]), min(dp[i-5], dp[i-6])))
		}
	}

	fmt.Println(dp)
	return dp[30]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

	// arr := []*Interval{
	// 	{0, 30},
	// 	{5, 10},
	// 	{15, 20},
	// 	{21, 22},
	// }
	//a := minMeetingRooms(arr)
	a := saanpseedi()
	fmt.Println(a)
}
