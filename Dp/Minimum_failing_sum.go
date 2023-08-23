package Dp

//Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
//
//A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right.
//Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
//[[100,-42,-46,-41],[31,97,10,-10],[-58,-51,82,89],[51,81,69,-51]]
//-36

func MinFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	dp[0] = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = matrix[0][i]
	}
	minSum := 0
	for i := 1; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			minSum = dp[i-1][j]
			if j-1 >= 0 && minSum > dp[i-1][j-1] {
				minSum = dp[i-1][j-1]
			}
			if j+1 < len(matrix[i]) && minSum > dp[i-1][j+1] {
				minSum = dp[i-1][j+1]
			}
			dp[i][j] = minSum + matrix[i][j]
		}
	}
	minSum = dp[len(matrix)-1][0]
	for i := 1; i < len(matrix[0]); i++ {
		if minSum > dp[len(matrix)-1][i] {
			minSum = dp[len(matrix)-1][i]
		}
	}
	return minSum
}
