package SlidingWindow

func MaxSumArray(nums []int, k int) int {
	sum := 0
	max_sum := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		sum = sum + nums[end]
		if end-start == k {
			if sum > max_sum {
				max_sum = sum
			}
			sum = sum - nums[start]
			start++
		}
	}
	return max_sum
}
