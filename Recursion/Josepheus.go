package Recursion

func GetPosition(n int, k int) int {
	if n == 1 {
		return 1
	}

	m := []int{}
	// 0 based indexing
	for i := 0; i < n; i++ {
		m = append(m, i)
	}

	start := 0
	for len(m) != 1 {
		out := (start + k - 1) % len(m)
		// delete element at index k
		m = append(m[:out], m[out+1:]...)
		start = out
	}

	return m[0] + 1
}
