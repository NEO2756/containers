/**
 * Definition of Interval:

 */

///////////////////////////////////////////
/*type char struct {
	frq int
	val rune
}
type myheap []char

func (h myheap) Len() int {
	return len(h)
}
func (h myheap) Swap(i, j int) {
	h[i].frq, h[j].frq = h[j].frq, h[i].frq
	h[i].val, h[j].val = h[j].val, h[i].val
}
func (h myheap) Less(i, j int) bool {
	return h[i].frq < h[j].frq
}
func (h *myheap) Push(x interface{}) {
	*h = append(*h, x.(char))
}
func (h *myheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *myheap) Top() char {
	return (*h)[0]
}

// func partitionLabels(s string) []int {

// 	f := make([]int, 26)
// 	ans := make([]int, 0)
// 	hp := myheap{}

// 	for i := 0; i < len(s); i++ {
// 		f[s[i]-'a']++
// 	}

// 	// for i := 0; i < len(s); i++ {
// 	// 	heap.Push(&hp, char{f[s[i]-'a'], string(s[i] - 'a')})
// 	// }

// 	start := 0
// 	heap.Push(&hp, char{f[s[0]-'a'], rune(s[0])})
// 	for i := 0; i < len(s); i++ {

// 		f[s[i]-'a']--

// 		//can we break it from here
// 		if hp.Len() > 0 {
// 			top := hp.Top()
// 			if f[top.val-'a'] == 0 {
// 				heap.Pop(&hp) //remove from heap
// 				if check(start, i, f) {
// 					start = i + 1
// 					ans = append(ans, i)
// 				}
// 			} else {
// 				heap.Push(&hp, char{f[s[i]-'a'], rune(s[i])})
// 			}
// 		} else {
// 			heap.Push(&hp, char{f[s[i]-'a'], rune(s[i])})
// 		}
// 	}
// 	return ans
// }

func partitionLabels(s string) []int {

	f := make([]int, 26)
	ans := make([]int, 0)

	for i := 0; i < len(s); i++ {
		f[s[i]-'a']++
	}

	start := 0
	for i := 0; i < len(s); i++ {

		f[s[i]-'a']--

		if f[s[i]-'a'] == 0 {
			if check(start, i, f, s) {
				ans = append(ans, i-start+1)
				start = i + 1
			}
		}
	}
	return ans
}

func check(start, end int, f []int, s string) bool {
	for i := start; i < end; i++ {
		if f[s[i]-'a'] > 0 {
			return false
		}
	}

	return true
}
*/
