package Graph

import (
	"math"
)

// Bellman_Ford_Algorithm Bellman_Ford_Algo is used for finding the shortest distance route from source to destination
// It will not work in case of negative weight cycle
func Bellman_Ford_Algorithm(edges []Edge, start int) {
	distances := make([]float64, len(edges))
	for i := 0; i < len(edges); i++ {
		if i != start {
			distances[i] = math.Inf(1)
		}
	}
	// We need to relax the vertices till edges-1
	for i := 0; i < len(edges); i++ {
		for _, e := range edges {
			//Relaxing each edge so that each vertex's distance will be minimum
			if distances[e.Src]+float64(e.Cost) < distances[e.Dest] {
				distances[e.Dest] = distances[e.Src] + float64(e.Cost)
			}
		}
	}
	// Checking for negative weight cycle
	for _, e := range edges {
		//Relaxing each edge so that each vertex's distance will be minimum
		if distances[e.Src]+float64(e.Cost) < distances[e.Dest] {
			distances[e.Dest] = distances[e.Src] + float64(e.Cost)
		}
	}
}

func Create_Graph_for_Bellman_Ford() {
	edge := []Edge{{0, 1, 5},
		{0, 2, 2},
		{1, 2, 1},
		{1, 3, 3},
		{2, 4, 4},
		{3, 4, 6}}
	Bellman_Ford_Algorithm(edge, 0)
}
