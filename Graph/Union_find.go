package Graph

import "fmt"

// Detect cycle in a graph


func IsCycle(g *Graph) bool {
   parent:=[]int{}
	for _, _ =range g.Vertices{
	   parent=append(parent,-1)
   }
	   for key,value:=range g.Vertices{
		   for k := range value.Vertices {
			   x:=find(parent,key)
			   y:=find(parent,k)
			  if x==y{
				  return false
			  }else{
				  Union(parent,key,k)
			  }
		   }
	   }
	return true
}

func Union(parent []int, x int ,y int){
   xset:=find(parent,x)
   yset:=find(parent,y)
   parent[xset]=yset
}

func find(parent []int, x int) int {
	if parent[x]==-1{
		return x
	}
	return find(parent,parent[x])
}
func Create_Graph(){
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
	fmt.Println(IsCycle(g))
}