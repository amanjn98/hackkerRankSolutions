package Graph

import (
	"math"
)

//(787)The Cheapest Flights Within K Stops
//There are n cities connected by some number of flights. You are given an array flights
//where flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi to city toi with cost pricei.
//You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops. If there is no such route, return -1.

// Example 1:
// Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
// Output: 700
// Explanation:
// The graph is shown above.
// The optimal path with at most 1 stop from city 0 to 3 is marked in red and has cost 100 + 600 = 700.
// Note that the path through cities [0,1,2,3] is cheaper but is invalid because it uses 2 stops.
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[src] = 0

	for i := 0; i <= k; i++ {
		tmp := make([]int, len(cost))
		copy(tmp, cost)
		for _, f := range flights {
			curr, nxt, price := f[0], f[1], f[2]
			if cost[curr] == math.MaxInt32 {
				continue
			}
			tmp[nxt] = min(tmp[nxt], cost[curr]+price)
		}
		cost = tmp
	}

	if cost[dst] == math.MaxInt32 {
		return -1
	}

	return cost[dst]
}

func Create_Flight_routes() int {
	var edges []Edge
	// Add new edges for the graph
	AddNewEdge(0, 1, 100, &edges)
	AddNewEdge(1, 2, 100, &edges)
	AddNewEdge(2, 0, 100, &edges)
	AddNewEdge(1, 3, 600, &edges)
	AddNewEdge(2, 3, 200, &edges)
	//return Cheapest_Flights(edges, 0, 3, 1)
	// [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
	flightdata := make([][]int, 5)
	//flightdata[0] = []int{0, 1, 1}
	//flightdata[1] = []int{0, 2, 5}
	//flightdata[2] = []int{1, 2, 1}
	//flightdata[3] = []int{2, 3, 1}
	flightdata[0] = []int{0, 1, 100}
	flightdata[1] = []int{1, 2, 100}
	flightdata[2] = []int{2, 0, 100}
	flightdata[3] = []int{1, 3, 600}
	flightdata[4] = []int{2, 3, 200}
	return findCheapestPrice(4, flightdata, 0, 3, 1)

}
