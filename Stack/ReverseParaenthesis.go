package Stack

import "fmt"

func ReverseParentheses(s string) string {
	stack := Stack{}
	fmt.Println(s)
	for _, c := range s {
		fmt.Println(c)
		if c == ')' {
			ele := ""
			for stack[len(stack)-1] != "(" {
				fmt.Println(ele)
				ele += stack.Pop()
			}
			stack.Pop()
			ele = reverse(ele)
			stack.Push(ele)
		} else {
			stack.Push(string(c))
		}
	}
	res := ""
	for len(stack) > 0 {
		x := stack.Pop()
		res = res + x
	}
	return reverse(res)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
