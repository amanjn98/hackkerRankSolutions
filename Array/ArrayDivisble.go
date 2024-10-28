package Array

/*
Check If Array Pairs Are Divisible by k
Given an array of integers arr of even length n and an integer k.

We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.

Return true If you can find a way to do that or false otherwise.

Example 1:

Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
Output: true
Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
Example 2:

Input: arr = [1,2,3,4,5,6], k = 7
Output: true
Explanation: Pairs are (1,6),(2,5) and(3,4).
Example 3:

Input: arr = [1,2,3,4,5,6], k = 10
Output: false
Explanation: You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.
*/
func CanArrange(arr []int, k int) bool {
	// Create an array to store the frequency of each remainder
	remainderCount := make([]int, k)

	// Iterate over the array and calculate the remainder of each number when divided by k
	for _, num := range arr {
		// Handle negative numbers by adjusting the remainder to be non-negative
		remainder := ((num % k) + k) % k
		remainderCount[remainder]++
	}

	// Now we need to check if we can form pairs such that the sum of each pair is divisible by k
	for r := 1; r <= k/2; r++ {
		// Special case for remainders equal to k/2 (like 3 in k=6)
		if r == k-r {
			// We must have an even number of elements with remainder r, to pair them together
			if remainderCount[r]%2 != 0 {
				return false
			}
		} else {
			// For any other remainder r, it must be paired with remainder k-r
			if remainderCount[r] != remainderCount[k-r] {
				return false
			}
		}
	}

	// Special case for remainder 0 (numbers divisible by k)
	if remainderCount[0]%2 != 0 {
		return false
	}

	// If all checks pass, return true
	return true
}
