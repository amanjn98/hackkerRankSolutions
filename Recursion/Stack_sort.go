package Recursion

import "fmt"

type SortStack struct {
	Items []int
}

func (s *SortStack) Push(ele int) {
	s.Items = append(s.Items, ele)
}

// Print() function
func (s *SortStack) Print() {
	length := len(s.Items)
	for i := 0; i < length; i++ {
		fmt.Print(s.Items[i], " ")
	}
	fmt.Println()
}
func (s *SortStack) Pop() int {
	length := len(s.Items)
	if length == 0 {
		return 0
	}
	res := s.Items[length-1]
	s.Items = s.Items[:length-1]
	return res
}

func Stack_Sort(s *SortStack) *SortStack {
	if len(s.Items) == 0 {
		return s
	}
	temp := s.Pop()
	Stack_Sort(s)
	return Arrange_Stack(s, temp)
}

func Arrange_Stack(s *SortStack, ele int) *SortStack {
	if len(s.Items) == 0 || s.Items[len(s.Items)-1] > ele {
		s.Push(ele)
		return s
	}
	temp := s.Pop()
	s = Arrange_Stack(s, ele)
	s.Push(temp)
	return s
}
