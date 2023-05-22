package Graph

import "math"

type GraphData struct {
	vertices map[string]map[string]int
}

func (g *GraphData) AddVertex(vertex string) {
	if _, ok := g.vertices[vertex]; !ok {
		g.vertices[vertex] = make(map[string]int)
	}
}

func (g *GraphData) AddEdge(from, to string, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)
	g.vertices[from][to] = weight
}

func Dijkstra(g GraphData, start string) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)
	for vertex := range g.vertices {
		distances[vertex] = math.MaxInt32
	}
	distances[start] = 0

	for i := 0; i < len(g.vertices); i++ {
		current := getMinVertex(distances, visited)
		visited[current] = true

		for neighbor, weight := range g.vertices[current] {
			newDist := distances[current] + weight
			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
			}
		}
	}

	return distances
}

func getMinVertex(distances map[string]int, visited map[string]bool) string {
	minDist := math.MaxInt32
	var minVertex string
	for vertex, dist := range distances {
		if !visited[vertex] && dist <= minDist {
			minDist = dist
			minVertex = vertex
		}
	}
	return minVertex
}

//g := Graph{vertices: make(map[string]map[string]int)}
//g.AddEdge("A", "B", 4)
//g.AddEdge("A", "C", 2)
//g.AddEdge("B", "C", 1)
//g.AddEdge("B", "D", 5)
//g.AddEdge("C", "D", 8)
//g.AddEdge("C", "E", 10)
//g.AddEdge("D", "E", 2)
//
//distances := Dijkstra(g, "A")
//fmt.Println(distances) // Output: map[A:0 B:4 C:2 D:9 E:11]
