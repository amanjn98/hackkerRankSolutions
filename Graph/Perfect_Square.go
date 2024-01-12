package Graph

import "math"

func NumSquares(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	count := 0
	queue := []int{n}
	visited := make(map[int]bool, int(math.Sqrt(float64(n)))+1)
	visited[n] = true
	for len(queue) != 0 {
		count++
		length := len(queue)
		for j := 1; j <= length; j++ {
			q := queue[0]
			queue = queue[1:]
			value := int(math.Sqrt(float64(q)))
			for i := 1; i <= value; i++ {
				if q-i*i == 0 {
					return count
				} else if visited[q-i*i] == false {
					queue = append(queue, q-i*i)
					visited[q-i*i] = true
				}
			}
		}
	}
	return count
}
