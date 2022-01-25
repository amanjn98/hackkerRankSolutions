//Leetcode problem
//784. Letter Case Permutation

package Recursion

import "strings"

func LetterCasePermutation(s string) []string {
	var opArray []string
	letterCaseRec(s,"", &opArray)
	return opArray
}

func letterCaseRec(s string, op string, i *[]string) {
	if len(s)==0{
		*i=append(*i,op)
		return
	}
	c:=s[0]
	if c >= '0' && c <= '9' {
		op=op+string(c)
		s=s[1:]
		letterCaseRec(s,op,i)
	}else{
		op1:=op+strings.ToLower(string(s[0]))

		op2 :=op+strings.ToUpper(string(s[0]))

		s=s[1:]
		letterCaseRec(s, op1, i)
		letterCaseRec(s, op2, i)
	}
	return
}

