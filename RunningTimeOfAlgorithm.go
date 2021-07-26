package main

import "fmt"

func runningTime(arr []int32) int32 {
	// Write your code here
	var countOp int32
   for i:=1;i<len(arr);i++{
   	j:=i-1
   	key:=arr[i]
   	for j>=0 && arr[j]>key{
   		countOp++
   		arr[j+1]=arr[j]
   		j--
	}
	arr[j+1]=key
   }
   return countOp
}

func main() {
	arr:=[]int32{2,1,3,1,2}
	fmt.Println(runningTime(arr))
}
