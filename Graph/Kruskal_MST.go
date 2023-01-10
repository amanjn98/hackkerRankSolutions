package Graph

import (
	"sort"
)

type Edge struct {
	Src		int
	Dest	int
	Cost 	int
}


func AddNewEdge(k1 int, k2 int, cost int, i *[]Edge) {
	var edge Edge
	edge.Src=k1
	edge.Dest=k2
	edge.Cost=cost
	*i=append(*i,edge)
}


func CreateMST(edges []Edge, nVertices int) int{
	 i:=0
	 var result []Edge
     for(nVertices< (nVertices -1)){
	  		x:=edges[i].Src
			y:=edges[i].Dest
			if x!=y{
				result=append(result,edges[i])
			}
			i++
	}
	i=0
	cost:=0
	for i<len(result) {
	  cost=cost+result[i].Cost
	  i++
	}
	return cost
}

func Kruskal_Mst(){
	var edges []Edge
	// Add new edges for the graph
	AddNewEdge(0,1,10,&edges)
	AddNewEdge(0,2,6,&edges)
	AddNewEdge(0,3,5,&edges)
	AddNewEdge(1,3,15,&edges)
	AddNewEdge(2,3,5,&edges)

	// Sort the array w.rt. cost of the edges
	sort.SliceStable(edges,func(i, j int) bool {
		return edges[i].Cost <edges[j].Cost
	})


}


