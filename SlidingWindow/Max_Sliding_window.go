package SlidingWindow

// https://leetcode.com/problems/sliding-window-maximum/description/
//You are given an array of integers nums, there is a sliding window of size k which is moving from the very
//left of the array to the very right. You can only see the k numbers in the window. Each time the sliding
//window moves right by one position. Return the max sliding window.
//
//
//Example 1:
//
//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
//Output: [3,3,5,5,6,7]
//Explanation:
//Window position                Max
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7

func MaxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var result = make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		// Remove all elements smaller than or equal to nums[i]
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		// Add i to queue
		queue = append(queue, i)

		// Skip reducing the window size if i is less than window size
		if i < k-1 {
			continue
		}

		// Delete all elements whose index is out of window range
		for len(queue) > 0 && queue[0] < i-k+1 {
			queue = queue[1:]
		}
		// First element is the largest element
		result[i-k+1] = nums[queue[0]]
	}
	return result
}
