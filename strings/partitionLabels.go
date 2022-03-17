func partitionLabels(s string) []int {

	f := make([]int, 26)
	m := make(map[byte]int)
	ans := make([]int, 0)

	for i := 0; i < len(s); i++ {
		f[s[i]-'a']++
		m[s[i]] = i
	}

	start := 0
	max := 0
	for i := 0; i < len(s); i++ {
		max = maxi(max, m[s[i]])
		f[s[i]-'a']--
		if f[s[i]-'a'] == 0 && m[s[i]] == max {
			ans = append(ans, i-start+1)
			start = i + 1
		}
	}
	return ans
}
func maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}