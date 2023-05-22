package Dp

func SubsetSum(sum int, arr []int) bool {
	dp := make([][]bool, len(arr)+1)
	for i := 0; i <= len(arr); i++ {
		dp[i] = make([]bool, sum+1)
		for j := 0; j <= sum; j++ {
			if i == 0 {
				dp[i][j] = false
			}
			if j == 0 {
				dp[i][j] = true
			}
		}
	}
	for i := 1; i <= len(arr); i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(arr)][sum]
}
