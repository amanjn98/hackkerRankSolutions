package Recursion

func DeleteMiddle(s *SortStack, k int) *SortStack {
	if s.Items[len(s.Items)-1] == k {
		s.Pop()
		return s
	}
	temp := s.Pop()
	s = DeleteMiddle(s, k)
	s.Push(temp)
	return s
}
