package Dp

import "math"

//function to get minimum number of coins to get a Sum
func Minimum_Num_Coins(sum int, arr []int) int {
	dp := make([][]int, len(arr)+1)
	for i := 0; i <= len(arr); i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			if i == 0 {
				dp[i][j] = math.MaxInt64 - 1
			} else if j == 0 {
				dp[i][j] = 0
			} else if i == 1 {
				dp[1][j] = j / arr[0]
			}
		}
	}

	for i := 2; i <= len(arr); i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] <= j {
				dp[i][j] = min(dp[i][j-arr[i-1]]+1, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(arr)][sum]
}
