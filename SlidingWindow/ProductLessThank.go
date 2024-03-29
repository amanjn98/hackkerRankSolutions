package SlidingWindow

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) < 1 {
		return 0
	}
	var start, end, count int
	product := nums[0]
	if product < k {
		count++
	}
	for end = 1; end < len(nums); end++ {
		product *= nums[end]
		if product < k {
			count += end - start + 1
		} else {
			for start < end && product >= k {
				product /= nums[start]
				start++
			}
			if product < k {
				count += end - start + 1
			}
		}
	}
	return count
}
