package Recursion

//Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]

func GenerateParenthesis(n int) []string {
	var opArray []string
	parenthesis("(", &opArray, n-1, n)
	return opArray
}

func parenthesis(op string, i *[]string, open int, closed int) {
	if open == 0 && closed == 0 {
		*i = append(*i, op)
		return
	}
	if open == 0 {
		addAllClose(op, i, closed)
		return
	} else if open >= closed {
		op = op + "("
		parenthesis(op, i, open-1, closed)
	} else {
		op1 := op + "("
		op2 := op + ")"
		parenthesis(op1, i, open-1, closed)
		parenthesis(op2, i, open, closed-1)
	}
	return
}

func addAllClose(op string, in *[]string, closed int) {
	for closed > 0 {
		op = op + ")"
		closed = closed - 1
	}
	*in = append(*in, op)
	return
}
