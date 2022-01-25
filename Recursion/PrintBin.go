package Recursion

//Print N-bit binary numbers having more 1s than 0s

func PrintNBinary(n int){
	op:=""
	PrintBinary(0, 0,op,n)
}

func PrintBinary(numOnes int, numZeros int,op string,n int){
	if n==0{
		println(op)
		return
	}

	if numOnes>numZeros{
		op1:=op + "0"
		op2:=op+"1"
		PrintBinary(numOnes, numZeros+1, op1,n-1)
		PrintBinary(numOnes+1,numZeros,op2,n-1)
	}else{
		op=op+"1"
		numOnes=numOnes+1
		PrintBinary(numOnes,numZeros,op,n-1)
	}

	return
}
//func addAllOne(op string, closed int) string {
//	for closed>0{
//		op=op+"1"
//		closed=closed-1
//	}
//	return op
//}
