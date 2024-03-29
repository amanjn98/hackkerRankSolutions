package Array

import "fmt"

// A set of numbers is given and the index by which the set needs to be shifted is specified. Example: input [5,7,8,4,9,1],
// index 3 -> [9,1,5,7,8,4] Ideally, make “in-place”, i.e. without creating an additional array
func RotateLeft(arr []int, count int) {
	//for i := 0; i < count; i++ {
	//	length := len(arr)
	//	last := arr[length-1]
	//	for j := length - 1; j > 0; j-- {
	//		arr[j] = arr[j-1]
	//	}
	//	arr[0] = last
	//}
	//fmt.Println(arr)
	p := 0
	for p <= count {
		last := arr[0]
		for i := 0; i < len(arr)-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[len(arr)-1] = last
		p++
	}
	fmt.Println(arr)
}
