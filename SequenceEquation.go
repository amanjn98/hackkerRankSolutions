package main

// Save the index in one of the array so that each index can be used every time needed
func permutationEquation(p []int32) []int32 {
	// Write your code here
	y := make([]int32, 0)
	for i := 0; i < len(p); i++ {
		y = append(y, 1+getIndex(int32(i)+1, p))
	}
	z := make([]int32, 0)
	for i := 0; i < len(p); i++ {
		z = append(z, y[y[i]-1])
	}
	return z
}

func getIndex(elem int32, a []int32) int32 {
	var i int32

	for i = 0; i < int32(len(a)); i++ {
		if elem == a[i] {
			return i
		}
	}

	return -1
}

// func main() {
// 	p:=[]int32{5,2,1,3,4}
// 	fmt.Println(permutationEquation(p))
// }
