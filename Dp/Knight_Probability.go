package Dp

//On an n x n chessboard, a knight starts at the cell (row, column) and attempts to make exactly k moves.
//The rows and columns are 0-indexed, so the top-left cell is (0, 0), and the bottom-right cell is (n - 1, n - 1).
//
//A chess knight has eight possible moves it can make, as illustrated below. Each move is two cells in a cardinal direction,
//then one cell in an orthogonal direction.
//
//
//Each time the knight is to move, it chooses one of eight possible moves uniformly at
//random (even if the piece would go off the chessboard) and moves there.
//
//The knight continues moving until it has made exactly k moves or has moved off the chessboard.
//
//Return the probability that the knight remains on the board after it has stopped moving.

func KnightProbability(n int, k int, row int, column int) float64 {
	moves := [][]int{[]int{1, 2}, []int{2, 1}, []int{-1, 2}, []int{2, -1}, []int{1, -2}, []int{-2, 1}, []int{-2, -1}, []int{-1, -2}}

	dp0 := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp0[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dp0[i][j] = 1.0
		}
	}
	var dp1 [][]float64
	for move := 1; move <= k; move++ {
		dp1 = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp1[i] = make([]float64, n)
			for j := 0; j < n; j++ {
				for _, s := range moves {
					is, js := i+s[0], j+s[1]
					if is >= 0 && is < n && js >= 0 && js < n {
						dp1[i][j] += 0.125 * dp0[is][js]
					}
				}
			}
		}
		dp0 = dp1
	}
	return dp0[row][column]
}
