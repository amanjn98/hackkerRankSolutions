package Recursion

import (
	"fmt"
)

func PermutationString(str string, l int, r int) {
	if l == r {
		fmt.Println(str)
	} else {
		for i := 1; i <= r; i++ {
			str = swap(str, i, l)
			PermutationString(str, l+1, r)
			str = swap(str, i, l)
		}
	}
}

func swap(str string, i, j int) string {
	strBytes := []byte(str)
	strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
	return string(strBytes)
}
