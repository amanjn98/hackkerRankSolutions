package Dp

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		nextUgly := minOf(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
		ugly[i] = nextUgly

		if nextUgly == ugly[i2]*2 {
			i2++
		}
		if nextUgly == ugly[i3]*3 {
			i3++
		}
		if nextUgly == ugly[i5]*5 {
			i5++
		}
	}
	return ugly[n-1]
}

func minOf(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
