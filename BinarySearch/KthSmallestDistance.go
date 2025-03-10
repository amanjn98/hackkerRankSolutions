package BinarySearch

import "sort"

//The distance of a pair of integers a and b is defined as the absolute difference between a and b.

//Given an integer array nums and an integer k, return the kth smallest distance among all the pairs nums[i] and nums[j] where 0 <= i < j < nums.length.

//Example 1:

//Input: nums = [1,3,1], k = 1
//Output: 0
//Explanation: Here are all the pairs:
//(1,3) -> 2
//(1,1) -> 0
//(3,1) -> 2
//Then the 1st smallest distance pair is (1,1), and its distance is 0.
//Example 2:

//Input: nums = [1,1,1], k = 2
//Output: 0
//Example 3:

//Input: nums = [1,6,1], k = 3
//Output: 5

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1]-nums[0]

	for low < high {
		mid := low + (high-low)/2
		if countPairs(nums, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func countPairs(nums []int, distance int) int {
	count, left := 0, 0
	for right := 1; right < len(nums); right++ {
		for nums[right]-nums[left] > distance {
			left++
		}
		count += right - left
	}
	return count
}
