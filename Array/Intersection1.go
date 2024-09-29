package Array

//Given two integer arrays nums1 and nums2, return an array of their intersection
//. Each element in the result must be unique and you may return the result in any order.
//
//
//
//Example 1:
//
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]
//Example 2:
//
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [9,4]
//Explanation: [4,9] is also accepted.

func intersection(nums1 []int, nums2 []int) []int {
	arr := make([]int, 1001)
	result := []int{}
	for i := 0; i < len(nums1); i++ {
		arr[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if arr[nums2[i]] > 0 {
			result = append(result, nums2[i])
			arr[nums2[i]] = 0
		}
	}
	return result
}
