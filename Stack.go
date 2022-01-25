package main

import "fmt"

type Stack struct {
	Items []int
	Top int
}

func (s *Stack) Push (ele int) {
	if len(s.Items)-1==s.Top{
		return
	}
	s.Items=append(s.Items,ele)
	s.Top++
}

func (s *Stack) Pop() int {
	if len(s.Items)==0{
		return 0
	}
	ele:=s.Items[s.Top-1]
	s.Top--
	return ele
}

func( s Stack) print(){
	fmt.Println(s.Items)
}

func main() {
	s:= Stack{}
	s.Push(5)
	s.Push(6)
	s.Push(8)
	s.print()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}