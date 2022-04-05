package Graph

import "fmt"

type Queue struct {
	Front int;
	Rear int;
	items []int;
}
func(q *Queue) push_queue(vertex int){
	if q.Front==-1{
		q.Front=0
	}
	q.items=append(q.items,vertex)
	q.Rear++
}

func(q *Queue) isEmptyQueue() int {
	if q.Front == -1 || q.Front > q.Rear {
		return 1;
	}
	return 0
}

func(q *Queue) popQueue() int {
	if q.isEmptyQueue()==1{
		fmt.Println("Queue Underflow")
		return -1
	}
	ele:=q.items[q.Front]
	q.Front++
	return ele
}
