package main

import (
	"fmt"
)

func pickingNumbers(a []int32) int32 {
	// Write your code here
    //number to be used for checking with other elem

    //res1:=0
	//for i:=0;i<len(a);i++{
    //  res:=1
    //  for j:=i+1;j<len(a);j++{
    //  	if a[j]==a[i]+1{
    //  		res=res+1
	//	}else if a[j]==a[i]-1{
	//		res=res+1
	//	}else if a[j]==a[i]{
	//		res=res+1
	//	}
	//  }
	//  if res>res1{
	//  	res1=res
	//  }
	//}
	//
    //return int32(res1)
	max:=0
	arr:=[100]int{0}
	for i:=0;i<len(a);i++{
		arr[a[i]]++
	}
	for i:=0;i<len(a)-1;i++{
		if arr[i]+arr[i+1]> max{
			max=arr[i]+arr[i+1]
		}
	}
	return int32(max)
}

func main() {
	a:=[]int32{1,2,2,3,1,2}
	fmt.Println(pickingNumbers(a))
}