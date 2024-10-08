package Graph

import "math"

// Second Minimum Time to Reach Destination
//A city is represented as a bi-directional connected graph with n vertices where each vertex is labeled from 1 to n (inclusive).
//The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi] denotes a bi-directional
//edge between vertex ui and vertex vi. Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.
//The time taken to traverse any edge is time minutes.
//
//Each vertex has a traffic signal which changes its color from green to red and vice versa every change minutes.
//All signals change at the same time. You can enter a vertex at any time, but can leave a vertex only when the signal is green.
//You cannot wait at a vertex if the signal is green.
//
//The second minimum value is defined as the smallest value strictly larger than the minimum value.
//
//For example the second minimum value of [2, 3, 4] is 3, and the second minimum value of [2, 2, 4] is 4.
//Given n, edges, time, and change, return the second minimum time it will take to go from vertex 1 to vertex n.

//Notes:
//
//You can go through any vertex any number of times, including 1 and n.
//You can assume that when the journey starts, all signals have just turned green.
//Input: n = 5, edges = [[1,2],[1,3],[1,4],[3,4],[4,5]], time = 3, change = 5
//Output: 13
//Explanation:
//The figure on the left shows the given graph.
//The blue path in the figure on the right is the minimum time path.
//The time taken is:
//- Start at 1, time elapsed=0
//- 1 -> 4: 3 minutes, time elapsed=3
//- 4 -> 5: 3 minutes, time elapsed=6
//Hence the minimum time needed is 6 minutes.
//
//The red path shows the path to get the second minimum time.
//- Start at 1, time elapsed=0
//- 1 -> 3: 3 minutes, time elapsed=3
//- 3 -> 4: 3 minutes, time elapsed=6
//- Wait at 4 for 4 minutes, time elapsed=10
//- 4 -> 5: 3 minutes, time elapsed=13
//Hence the second minimum time is 13 minutes.

// Input: n = 2, edges = [[1,2]], time = 3, change = 2
//Output: 11
//Explanation:
//The minimum time path is 1 -> 2 with time = 3 minutes.
//The second minimum time path is 1 -> 2 -> 1 -> 2 with time = 11 minutes.

type Pair struct {
	cost int
	node int
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	graph := make(map[int][]int)
	for _, edg := range edges {
		graph[edg[0]] = append(graph[edg[0]], edg[1])
		graph[edg[1]] = append(graph[edg[1]], edg[0])
	}
	dist := make([]int, n+1)
	dist1 := make([]int, n+1)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]][edges[i][1]] = 1
		graph[edges[i][1]][edges[i][0]] = 1
	}
	for i := 0; i < n+1; i++ {
		dist[i] = math.MaxInt32
		dist1[i] = math.MaxInt32
	}
	queue := []Pair{Pair{cost: 0, node: 1}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node.node] {
			newcost := time + node.cost
			if (node.cost)/change%2 == 1 {
				newcost += change - (node.cost)%change
			}
			if dist[neighbor] > newcost {
				dist1[neighbor] = dist[neighbor]
				dist[neighbor] = newcost
			} else if dist[neighbor] < newcost && dist1[neighbor] > newcost {
				dist1[neighbor] = newcost
				queue = append(queue, Pair{cost: newcost, node: neighbor})
			}
		}
	}
	return dist[n]
}
