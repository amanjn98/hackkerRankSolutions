package Linked_list

import "fmt"

// Implementation of queue using Linked list

type Linked struct {
	data int
	next *Linked
}

type Queue struct {
	Rear  *Linked
	Front *Linked
}

func (q *Queue) add(key int) {
	temp := Linked{
		data: key,
		next: nil,
	}
	if q.Rear == nil {
		q.Rear = &temp
		q.Front = &temp
		return
	}
	q.Rear.next = &temp
	q.Rear = &temp
}

func (q *Queue) delete() int {
	if q.Front == nil {
		return -1
	}
	temp := q.Front.data
	q.Front = q.Front.next
	if q.Front == nil {
		q.Rear = nil
	}
	return temp
}
func (q *Queue) Print() {
	if q.Front == nil {
		return
	}
	value := q.delete()
	for value != -1 {
		fmt.Println(value)
		value = q.delete()
	}
}

func QueueImp() {
	q := Queue{}
	q.add(5)
	q.add(20)
	q.add(25)
	q.Print()
}
