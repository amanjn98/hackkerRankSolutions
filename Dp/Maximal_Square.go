package Dp

//Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
//Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
//Output: 4
func MaximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	max := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
			}
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}

	}
	row := []int{0, -1, -1}
	col := []int{-1, 0, -1}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				if dp[row[0]+i][col[0]+j] == 0 || dp[row[1]+i][col[1]+j] == 0 || dp[row[2]+i][col[2]+j] == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[row[0]+i][col[0]+j], dp[row[1]+i][col[1]+j])
					dp[i][j] = min(dp[i][j], dp[row[2]+i][col[2]+j]) + 1
				}

			} else {
				dp[i][j] = 0
			}
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}

	}
	return max * max
}

func CalculateMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
