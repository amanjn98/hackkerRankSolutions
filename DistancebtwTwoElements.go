package main

import "fmt"

func distance(arr []int,n int, x int, y int) int {
	dist:=-1;
	min:=9999999;
	for i:=0;i<n;i++{
		if x==arr[i] || y==arr[i] {
			if   dist!=-1 && arr[dist] != arr[i]{
				dist=i-dist;
				if min>dist  {
					min=dist;
				}
			}
			dist=i;
		}
	}
	if min==9999999 {
		return -1;
	}else{
		return min;
	}
}

func main() {
	a:=[]int{96, 82, 48, 76 ,34 ,19, 7, 80 ,36 ,57 ,77, 34 ,19 ,35, 5, 57, 16, 66, 42 ,62 ,89 ,19 ,60, 19, 25, 16, 20, 51,
		77, 75, 12, 73, 8, 11, 100, 93, 81, 58, 72, 17, 14, 48, 2, 33, 82, 6, 41, 49, 72, 34, 10, 12, 53, 21, 30, 77, 36, 49, 79, 13, 75, 42}
	fmt.Println(distance(a,len(a),36,33))
}