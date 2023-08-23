package Recursion

import "fmt"

//Print N bits binary number 's prefix having more than 1 than 0s
//Input:  N = 3
//Output: 111 110 101
//Explanation: 111, 110 and 101 have more
//than or equal 1's than 0's

func PrintNBit(n int) {
	if n%2 == 0 {
		PrintBinaryRec(n/2, n/2, "")
	} else {
		PrintBinaryRec(n/2+1, n/2, "")
	}
}

func PrintBinaryRec(numOnes int, numZeros int, op string) {
	if numOnes == 0 && numZeros == 0 {
		return
	}
	if numOnes >= numZeros {
		op := op + "1"
		fmt.Println(op)
		numOnes = numOnes - 1
		PrintBinaryRec(numOnes, numZeros, op)
	} else {
		op := op + "0"
		fmt.Println(op)
		numZeros = numZeros - 1
		PrintBinaryRec(numOnes, numZeros, op)
	}
	return
}