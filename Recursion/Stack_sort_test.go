package Recursion

import (
	"testing"
)

func TestStackSort(t *testing.T) {
	s := SortStack{}
	s.Push(3)
	s.Push(7)
	s.Push(5)
	s.Push(3)
	Stack_Sort(&s)
	s.Print()
	s.Pop()
	s.Print()
}
