package Dp

import "math"

//Given a triangle array, return the minimum path sum from top to bottom.
//
//For each step, you may move to an adjacent number of the row below. More formally,
//if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
func MinimumTotal(triangle [][]int) int {

	dp := make([][]int, len(triangle))
	dp[0] = make([]int, len(triangle[0]))
	dp[0][0] = triangle[0][0]
	minSum := math.MaxInt64
	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		minSum := math.MaxInt64
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 >= 0 && minSum > dp[i-1][j-1] {
				minSum = dp[i-1][j-1]
			}
			if j < len(triangle[i-1]) && minSum > dp[i-1][j] {
				minSum = dp[i-1][j]
			}
			dp[i][j] = minSum + triangle[i][j]
		}
	}
	minSum = dp[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if minSum > dp[len(triangle)-1][i] {
			minSum = dp[len(triangle)-1][i]
		}
	}
	return minSum
}
