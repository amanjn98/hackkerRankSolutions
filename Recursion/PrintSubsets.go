package Recursion

import "fmt"

func PrintSubset(in string, op string) {
	if len(in) == 0 {
		fmt.Println(op)
		return
	}
	PrintSubset(in[1:], op)
	PrintSubset(in[1:], op+string(in[0]))
}
