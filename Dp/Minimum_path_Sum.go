package Dp

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
// leetcode : Minimum path sum

//Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
//Output: 7
//Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

func CalculateSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = grid[0][0]
	// Only way to reach [1][0] is the cumulative sum of [0][0] and [1][0] and we can only move right for this path
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// Only way to reach [0][1] is the cumulative sum of [0][0] and [][1] and we can only move down for this path
	for j := 1; j < len(grid); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func isValid(x int, y int, M int, N int) bool {
	return (x >= 0 && x < M) && (y >= 0 && y < N)
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
