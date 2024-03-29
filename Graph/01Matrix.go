package Graph

import "math"

//Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
//
//The distance between two adjacent cells is 1.


func UpdateMatrix(mat [][]int) [][]int {
	var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(mat), len(mat[0])
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
		        queue = append(queue, [2]int{i, j})
			} else {
		         mat[i][j] = math.MaxInt64
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
		x, y := cur[0]+dir[0], cur[1]+dir[1]
		if x < 0 || x >= m || y < 0 || y >= n || mat[x][y] <= mat[cur[0]][cur[1]] {
		continue
	}
		queue = append(queue, [2]int{x, y})
		mat[x][y] = mat[cur[0]][cur[1]] + 1
	}
	}

		return mat
}