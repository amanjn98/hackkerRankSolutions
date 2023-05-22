package Recursion

import "fmt"

//I/p:"ABC"
//O/p : "A B C" ,"A BC","ABC","AB C
func PermuationWithSpaces(in string, op string) {
	if len(in) == 0 {
		fmt.Println(op)
		return
	}
	op1 := op + " " + string(in[0])
	op2 := op + string(in[0])
	PermuationWithSpaces(in[1:], op1)
	PermuationWithSpaces(in[1:], op2)
}
