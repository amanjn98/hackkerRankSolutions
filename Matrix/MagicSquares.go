package Matrix

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	count := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[i])-3; j++ {
			if checkMagicSquare(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func checkMagicSquare(grid [][]int, x int, y int) bool {
	distinctNo := make([]bool, 10)
	rowSum := make([]int, 3)
	colSum := make([]int, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rowSum[i] += grid[x+i][y+j]
			colSum[j] += grid[x+i][y+j]
			if grid[x+i][y+j] <= 0 || grid[x+i][y+j] > 9 || distinctNo[grid[x+i][y+j]] == true {
				return false
			}
			distinctNo[grid[x+i][y+j]] = true
		}
	}
	for i := 1; i < 3; i++ {
		if rowSum[i] != colSum[i] || rowSum[i] != rowSum[i-1] || colSum[i] != colSum[i-1] {
			return false
		}
	}
	digSum1 := grid[x][y] + grid[x+1][y+1] + grid[x+2][y+2]
	digSum2 := grid[x][y+2] + grid[x+1][y+1] + grid[x+2][y]
	if digSum1 != rowSum[0] || digSum1 != digSum2 {
		return false
	}
	return true
}
