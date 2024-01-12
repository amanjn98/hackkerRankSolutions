package Array

import "fmt"

// Postfix to infix conversion

type Stack struct {
	Items []string
	count int
}

func (s *Stack) Push(str string) {
	s.Items = append(s.Items, str)
	s.count++
}

func (s *Stack) Pop() string {
	if len(s.Items) == 0 {
		return ""
	}
	temp := s.Items[s.count-1]
	s.Items = s.Items[:s.count-1]
	s.count--
	return temp
}
func checkOperator(op rune) bool {
	if op >= '!' && op <= '@' {
		return true
	}
	return false
}

func checkOperand(op rune) bool {
	if (op >= 'a' && op <= 'z') || (op >= 'A' && op <= 'Z') {
		return true
	}
	return false
}
func Postfix_to_Infix(s string) {
	stack := Stack{
		Items: []string{},
		count: 0,
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if checkOperator(rune(s[i])) {
			op1 := stack.Pop()
			op2 := stack.Pop()
			res = "(" + op2 + string(rune(s[i])) + op1 + ")"
			stack.Push(res)
		} else {
			stack.Push(string(rune(s[i])))
		}
	}
	fmt.Println(res)
}
