package Graph

import "fmt"

type Node struct {
	Key      int;
	Vertices map[int]*Node;
}
type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Node

	directed bool
}

// NewVertex We then create a constructor function for the Key
func NewVertex(key int) *Node {
	return &Node{
		Key:      key,
		Vertices: map[int]*Node{},
	}
}
func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}
func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Node{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Node{},
	}
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	//if !g.directed && v1.Key != v2.Key {
	//	v2.Vertices[v1.Key] = v1
	//}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func BFS_Traversal(){

}
func create_graph() *Graph {
	g:=NewUndirectedGraph()
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	return g
}
func BFS(){
	g:=create_graph()
	fmt.Println(g)
}