package Array

//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
//You must write an algorithm that runs in O(n) time and without using the division operation.

//Example 1:
//
//Input: nums = [1,2,3,4]
//Output: [24,12,8,6]
//Example 2:
//
//Input: nums = [-1,1,0,-3,3]
//Output: [0,0,9,0,0]
//
//
//Constraints:
//
//2 <= nums.length <= 105
//-30 <= nums[i] <= 30
//The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer

// First take the left product excluding itself
// Then take the right product from the end and multiplying the left product
// Finally take the product of both and store it in the new product
func productExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums))
	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = nums[i-1] * leftProduct[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		leftProduct[i] = right * leftProduct[i]
		right *= nums[i]
	}
	return leftProduct
}
