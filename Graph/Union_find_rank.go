package Graph

import "fmt"

type NodeForDisjoint struct {
	rank int
	parent int
}

func CreateGraphForRank(){
	g:=NewUndirectedGraph()
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(3, 4)
	fmt.Println(IscycleDisjoint(g))
}

func findOperation(subset []NodeForDisjoint, i int) int {
	if subset[i].parent !=i{
		findOperation(subset,subset[subset[i].parent].parent)
	}
	return subset[i].parent
}

func unionOperation(subset []NodeForDisjoint, x int , y int){
	if subset[x].rank <subset[y].rank{
		subset[x].parent=y
	}else if subset[x].rank >subset[y].rank{
		subset[y].rank=x
	}else{
		subset[x].parent=y
		subset[y].rank++
	}
}

func IscycleDisjoint(g *Graph) bool{
	var node []NodeForDisjoint
	for key,_:=range g.Vertices{
		var data NodeForDisjoint
		data.rank=0
		data.parent=key
		node=append(node,data)
	}

	for key,value:=range g.Vertices{
		for k,_ := range value.Vertices{
		x:=node[key].parent
		y:=node[k].parent
		 if findOperation(node,x)==findOperation(node,y){
	        return true
		 }else{
			 unionOperation(node,x,y)
		 }
		}
	}
return false
}