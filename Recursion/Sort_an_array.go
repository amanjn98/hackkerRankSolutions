package Recursion

func Sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	temp := arr[len(arr)-1]
	arr = arr[0 : len(arr)-1]
	Sort(arr)
	return Arrange(arr, temp)
}

func Arrange(arr []int, ele int) []int {
	if len(arr) == 0 || ele > arr[len(arr)-1] {
		arr = append(arr, ele)
		return arr
	}
	temp := arr[len(arr)-1]
	arr = arr[0 : len(arr)-1]
	arr = Arrange(arr, ele)
	arr = append(arr, temp)
	return arr
}
