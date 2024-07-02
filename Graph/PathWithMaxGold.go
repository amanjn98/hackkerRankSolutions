package Graph

// 1219. Path with Maximum Gold
// In a gold mine grid of size m x n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.
//
// Return the maximum amount of gold you can collect under the conditions:
//
// Every time you are located in a cell you will collect all the gold in that cell.
// From your position, you can walk one step to the left, right, up, or down.
// You can't visit the same cell more than once.
// Never visit a cell with 0 gold.
// You can start and stop collecting gold from any position in the grid that has some gold.
// Dfs+Backtracking
func GetMaximumGold(grid [][]int) int {
	maxSum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				gold := backTrackForGold(grid, i, j)
				if maxSum < gold {
					maxSum = gold
				}
			}
		}
	}
	return maxSum
}

func backTrackForGold(grid [][]int, row, col int) int {
	m := len(grid)
	n := len(grid[0])
	maxGold := 0
	rowArray := []int{0, -1, 0, 1}
	colArray := []int{-1, 0, 1, 0}
	currGold := grid[row][col]
	grid[row][col] = 0
	for i := 0; i < 4; i++ {
		if isValidCheck(row+rowArray[i], col+colArray[i], m, n, grid) {
			maxGold = Max(maxGold, currGold+backTrackForGold(grid, row+rowArray[i], col+colArray[i]))
		} else {
			maxGold = Max(maxGold, currGold) // it is required since it will return directly 0 instead of currentValue
		}
	}
	grid[row][col] = currGold
	return maxGold
}

func isValidCheck(row int, col int, M int, N int, grid [][]int) bool {
	if row < 0 || row >= M || col < 0 || col >= N || grid[row][col] == 0 {
		return false
	}
	return true
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
