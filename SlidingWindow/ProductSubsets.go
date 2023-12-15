package SlidingWindow

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	left := 0
	current := 1
	ans := 0

	for right, v := range nums {
		current *= v
		for current >= k {
			current = current / nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
