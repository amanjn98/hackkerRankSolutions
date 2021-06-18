package main

import "fmt"

/*
 * Complete the 'angryProfessor' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY a
 */

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	count:=0
	for i:=0;i<len(a);i++{
		if a[i]<=0{
			count++
		}
		if int32(count)==k{
			return "NO"
		}
	}


	return "YES"
}

func main()  {
	a:=[]int32{-1,-1,0,2,3}
	fmt.Println(angryProfessor(4,a))
}

