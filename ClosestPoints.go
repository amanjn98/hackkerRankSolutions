package main

//K Closest Points to Origin


func kClosest(points [][]int, k int) [][]int {

	distance:=[]int{}
	for i:=0;i<len(points);i++{
		temp:=0
		for j:=0;j<len(points);j++{
			temp=points[i][j]*points[i][j]+temp
			distance=append(distance,temp)
		}
	}
	return nil
}