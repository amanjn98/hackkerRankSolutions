package Graph

//Find the shortest distance of every cell from a landmine inside a maze

//Given a maze in the form of a rectangular matrix, filled with either O, X, or M, where O represents an open cell,
//X represents a blocked cell, and M represents landmines in the maze, find the shortest distance of every open cell in the maze from its nearest mine.
//We are only allowed to travel either of the four directions, and diagonal moves are not allowed. We can assume cells with the mine have distance 0.
//Also, blocked and unreachable cells have distance -1.

//For example,
//
//Input: 6 × 5 matrix filled with O (Open cell), X (Blocked Cell), and M (Landmine).
//
//O   M   O   O   X
//O   X   X   O   M
//O   O   O   O   O
//O   X   X   X   O
//O   O   M   O   O
//O   X   X   M   O
//
//Output:
//
//1   0   1   2  -1
//2  -1  -1   1   0
//3   4   3   2   1
//3  -1  -1  -1   2
//2   1   0   1   2
//3  -1  -1   0   1

type NodeData struct {
	X        int
	Y        int
	Distance int
}

func verify(row int, col int, M int, N int) bool {
	return (row >= 0 && row < M) && (col >= 0 && col < N)
}

func isSafe(matrixValue rune, dist int) bool {
	return matrixValue == 'O' && dist == -1
}

func CalculateDistance(matrix [][]rune) [][]int {
	result := make([][]int, len(matrix))
	queue := make([]NodeData, 0)
	M := len(matrix)
	N := len(matrix[0])
	for i := 0; i < M; i++ {
		result[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if matrix[i][j] == 'M' {
				var node NodeData
				node.X = i
				node.Y = j
				node.Distance = 0
				queue = append(queue, node)
			} else {
				result[i][j] = -1
			}
		}
	}
	row := []int{0, -1, 0, 1}
	col := []int{-1, 0, 1, 0}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			if verify(row[i]+node.X, col[i]+node.Y, M, N) &&
				isSafe(matrix[row[i]+node.X][col[i]+node.Y], result[row[i]+node.X][col[i]+node.Y]) {
				result[row[i]+node.X][col[i]+node.Y] = node.Distance + 1
				queue = append(queue, NodeData{node.X + row[i], node.Y + col[i], node.Distance + 1})
			}
		}
	}
	return result
}
