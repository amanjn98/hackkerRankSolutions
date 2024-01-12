package Array

import "math"

//Given a signed 32-bit integer x, return x with its digits reversed.
//If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

func reverse(x int) int {
	rem := 0
	temp := x
	num := []int{}

	for temp != 0 {
		rem = temp % 10
		temp = temp / 10
		num = append(num, rem)
	}
	res := 0
	for i := 0; i < len(num); i++ {
		res += num[i] * int(math.Pow(float64(10), float64(len(num)-1-i)))
	}
	if res > math.MaxInt32 {
		return 0
	}
	if res < math.MinInt32 {
		return 0
	}
	return res
}
