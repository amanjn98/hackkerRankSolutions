package Dp

func CountSubsetSum(sum int, arr []int) int {
	dp := make([][]int, len(arr)+1)
	for i := 0; i <= len(arr); i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			if i == 0 {
				dp[i][j] = 0
			}
			if j == 0 {
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i <= len(arr); i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] <= j {
				dp[i][j] = 1 + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(arr)][sum]
}
