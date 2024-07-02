package Array

import "sort"

//You are given a sorted integer array arr containing 1 and prime numbers, where all the integers of arr are unique.
//You are also given an integer k. For every i and j where 0 <= i < j < arr.length, we consider the fraction arr[i] / arr[j].
//
//Return the kth smallest fraction considered. Return your answer as an array of integers of size 2, where answer[0] == arr[i] and answer[1] == arr[j].

func kthSmallestPrimeFraction(arr []int, k int) []int {
	mp := make(map[float64][]int)
	nums := make([]float64, 0)
	c := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			nums = append(nums, float64(arr[i])/float64(arr[j]))
			mp[nums[c]] = []int{arr[i], arr[j]}
			c++
		}
	}
	sort.Float64s(nums)
	return mp[nums[k-1]]
}
