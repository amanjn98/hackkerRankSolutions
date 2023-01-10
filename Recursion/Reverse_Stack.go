package Recursion

func ReverseStack(s *SortStack) *SortStack {
	if len(s.Items) == 0 {
		return s
	}
	temp := s.Pop()
	s = ReverseStack(s)
	s = insert(s, temp)
	return s
}

func insert(s *SortStack, ele int) *SortStack {
	if len(s.Items) == 0 {
		s.Push(ele)
		return s
	}
	temp := s.Pop()
	insert(s, ele)
	s.Push(temp)
	return s
}
