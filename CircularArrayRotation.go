package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	arr:=make([]int32,len(a))
	for i:=0;i<len(a);i++{
		j:=(int32(i)+k)%int32(len(a))
		arr[j]=a[i]
	}
	for i:=0;i<len(queries);i++{
		queries[i]=arr[queries[i]]
	}
	return queries
}

func main(){
	a:=[]int32{1,2,3}
	queries:=[]int32{0,1,2}
	fmt.Println(circularArrayRotation(a,2,queries))


}
