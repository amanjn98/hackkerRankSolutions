package Graph

//Find the shortest path from source to destination
//
//Given a chessboard, find the shortest distance (minimum number of steps) taken by a knight to reach a given destination from a given source.
//
//For example,
//
//Input:
//
//N = 8 (8 × 8 board)
//Source = (7, 0)
//Destination = (0, 7)
//
//Output: Minimum number of steps required is 6

type NodeDataForChess struct {
	x    int
	y    int
	dist int
}

func isValid(x int, y int, M int, N int) bool {
	return (x >= 0 && x < M) && (y >= 0 && y < N)
}

func CalculateDistanceForChess(x int, y int, N int, destX int, destY int) int {
	visited := make(map[NodeDataForChess]bool)
	var node NodeDataForChess
	node.x = x
	node.y = y
	node.dist = 0
	queue := []NodeDataForChess{}
	queue = append(queue, node)
	row := []int{2, 2, -2, -2, 1, 1, -1, -1}
	col := []int{-1, 1, 1, -1, 2, -2, 2, -2}
	m := N
	n := N
	for len(queue) != 0 {
		nodeData := queue[0]
		queue = queue[1:]
		if !visited[nodeData] {
			visited[nodeData] = true
			if destX == nodeData.x && destY == nodeData.y {
				return nodeData.dist
			}
			for i := 0; i < 8; i++ {
				X := nodeData.x + row[i]
				Y := nodeData.y + col[i]
				if isValid(X, Y, m, n) {
					node.x = X
					node.y = Y
					node.dist = nodeData.dist + 1
					queue = append(queue, node)
				}
			}
		}
	}

	return -1
}
