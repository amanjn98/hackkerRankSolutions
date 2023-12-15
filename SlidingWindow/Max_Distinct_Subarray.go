package SlidingWindow

//You are given an integer array nums and an integer k. Find the maximum subarray sum of all the subarrays of nums that meet the following conditions:
//
//The length of the subarray is k, and
//All the elements of the subarray are distinct.
//Return the maximum subarray sum of all the subarrays that meet the conditions. If no subarray meets the conditions, return 0.
//
//A subarray is a contiguous non-empty sequence of elements within an array.
//Input: nums = [1,5,4,2,9,9,9], k = 3
//Output: 15
//Explanation: The subarrays of nums with length 3 are:
//- [1,5,4] which meets the requirements and has a sum of 10.
//- [5,4,2] which meets the requirements and has a sum of 11.
//- [4,2,9] which meets the requirements and has a sum of 15.
//- [2,9,9] which does not meet the requirements because the element 9 is repeated.
//- [9,9,9] which does not meet the requirements because the element 9 is repeated.
//We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions

func MaximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if k <= 0 || k > n {
		return 0
	}
	frequency := make(map[int]int)
	start, sum, max_sum := 0, 0, 0
	//for loop till end reaches n
	for end := 0; end < n; end++ {
		frequency[nums[end]]++
		sum += nums[end]
		if end-start+1 == k {
			//distinct elements condition
			if len(frequency) == k {
				max_sum = max(max_sum, sum)
			}
			if frequency[nums[start]] == 1 {
				delete(frequency, nums[start])
			} else {
				frequency[nums[start]]--
			}
			sum -= nums[start]
			start++
		}
	}
	return int64(max_sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
