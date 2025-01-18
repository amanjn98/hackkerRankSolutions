package Array

//Given an integer array nums, return the number of subarrays filled with 0.
//
//A subarray is a contiguous non-empty sequence of elements within an array.
//
//
//
//Example 1:
//
//Input: nums = [1,3,0,0,2,0,0,4]
//Output: 6
//Explanation:
//There are 4 occurrences of [0] as a subarray.
//There are 2 occurrences of [0,0] as a subarray.
//There is no occurrence of a subarray with a size more than 2 filled with 0. Therefore, we return 6.
//Example 2:
//
//Input: nums = [0,0,0,2,0,0]
//Output: 9
//Explanation:
//There are 5 occurrences of [0] as a subarray.
//There are 3 occurrences of [0,0] as a subarray.
//There is 1 occurrence of [0,0,0] as a subarray.
//There is no occurrence of a subarray with a size more than 3 filled with 0. Therefore, we return 9.
//Example 3:
//
//Input: nums = [2,10,2019]
//Output: 0
//Explanation: There is no subarray filled with 0. Therefore, we return 0.

//Let's think of a sliding window of size k that slides from the start of the subarray S until the right edge of the window hits the end of S.
//
//When k = 1, the start of the sliding window can be placed over each element of S.
//[1] 2  3  4
// 1 [2] 3  4
// 1  2 [3] 4
// 1  2  3 [4]
//Thus we have n distinct subarrays of size 1.
//
//When k = 2 , the start of the sliding window can be placed over each element of S except the last.
//[1  2] 3  4
// 1 [2  3] 4
// 1  2 [3  4]
//So we have n - 1 distinct subarrays of size 2.
//
//And so on.
//
//Finally,
//
//When k = n, the sliding window covers the entirety of S.
//[1  2   3  4]
//We have 1 subarray of size n.
//
//Hence, the total number of distinct subarrays that can be formed out of S is:
//
//n + (n - 1) + (n - 2) + ... + 1 =
//((n + 1) * n) / 2

func zeroFilledSubarray(nums []int) int64 {
	ans := 0
	count := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			count++
		} else if count > 0 {
			ans += calculate(count)
			count = 0
		}
	}
	if count != 0 {
		ans += calculate(count)
	}
	return int64(ans)
}

func calculate(n int) int {
	return (n * (n + 1)) / 2
}
