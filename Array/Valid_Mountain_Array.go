package Array

//leetcode problem Valid Mountain array
func validMountainArray(arr []int) bool {
	found:=false
	if len(arr)<3{
		return false
	}
	i:=0
	// this loop check whether it is strictly increasing or not
	for i=0;i<len(arr)-1;i++{
		if arr[i]==arr[i+1]{
			return false
		}
		if arr[i]>arr[i+1]{
			found=false
			break
		}
		found=true
	}
	// if it is only strictly increasing or decreasing
	// [9,8,7,6,5,4] or [1,2,3,4,5]
	if i==len(arr)-1 || i==0{
		return false
	}
	// this loop check whether it is stricly decreasing or not
	for j:=i+1;j<len(arr);j++{
		if arr[j]==arr[j-1]{
			return false
		}
		if arr[j-1]<arr[j]{
			found=false
			break
		}
		found=true

	}
	return found
}
