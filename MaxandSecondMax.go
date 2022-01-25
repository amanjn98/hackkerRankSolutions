package main

func largestAndSecondLargest(arr []int) int {
	min:=arr[0]
	first:=arr[0]
	for _,n:=range arr{
		if n>first{
			first=n
		}
		if min>n{
			min=n
		}
	}
	second:=min

	for _,n:=range arr{
		if n>=second && n<first{
			second=n
		}
	}

	if second==first{
		return -1
	}
	return second
}
