package Heap

//k-closest-points-to-origin

func KClosest(points [][]int, k int) [][]int {
	if len(points)==0{
		return [][]int{}
	}

	var heap [][]int
	for _,pair:=range points{
		x,y:=pair[0],pair[1]
		dist:=x*x+y*y
		heap=append(heap,[]int{dist,x,y})
	}
	buildheap(heap)
	var result [][]int
	for k>0{
		result=append(result,[]int{heap[0][1],heap[0][2]})
		pop(&heap)
		k--
	}
	return result
}

func pop(heap *[][]int){
	n := len(*heap)
	lastElement := (*heap)[n - 1]
	(*heap)[0] = lastElement
	*heap = (*heap)[:n-1]
	heapify(heap, 0)
}
func buildheap(heap [][]int){
	n:=len(heap)
	startIdx:=n/2-1
	for i:=startIdx;i>=0;i--{
		heapify(&heap,i)
	}
}

func heapify(heap *[][]int,i int){
	smallest:=i
	lchild:=2*i+1
	rchild:=2*i+2

	if lchild<len(*heap) && (*heap)[lchild][0]<(*heap)[smallest][0]{
		smallest=lchild
	}

	if rchild<len(*heap) && (*heap)[rchild][0]<(*heap)[smallest][0]{
		smallest=rchild
	}

	if smallest!=i{
		temp:=(*heap)[smallest]
		(*heap)[smallest]=(*heap)[i]
		(*heap)[i]=temp
		heapify(heap,smallest)
	}
}