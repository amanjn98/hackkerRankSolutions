package Array

import (
	"fmt"
	"sort"
)

//Given an integer array nums that does not contain any zeros, find the largest positive integer k such that -k also exists in the array.
//
//Return the positive integer k. If there is no such integer, return -1.

//Example 1:
//
//Input: nums = [-1,2,-3,3]
//Output: 3
//Explanation: 3 is the only valid k we can find in the array.

func FindMaxK(nums []int) int {
	sort.Ints(nums)
	maxInt := -1
	fmt.Println(nums)
	j := len(nums) - 1
	i := 0
	for i < len(nums)-1 && j > 0 {
		if -1*nums[i] > nums[j] {
			i++
		} else if -1*nums[i] < nums[j] {
			j--
		} else if -1*nums[i] == nums[j] {
			if maxInt < nums[j] {
				maxInt = nums[j]
			}
			i++
			j--
		}
	}
	return maxInt
}
