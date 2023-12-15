package Recursion

import "fmt"

func GenerateSubsetsWithProductLessThanOrEqual(arr []int, k int) int {
	subset := []int{}
	var count int
	GenerateSubsets(arr, k, 0, subset, &count)
	return count
}

func GenerateSubsets(arr []int, k, index int, subset []int, i *int) {
	if index == len(arr) {
		// Check if the product of subset elements is less than or equal to k
		product := 1
		for _, num := range subset {
			product *= num
		}
		if product <= k {
			// Append a copy of the subset to the result
			temp := make([]int, len(subset))
			copy(temp, subset)
			*i = *i + 1
		}
		return
	}

	// Include the current element in the subset and recurse
	subset = append(subset, arr[index])
	GenerateSubsets(arr, k, index+1, subset, i)

	// Exclude the current element from the subset and recurse
	subset = subset[:len(subset)-1]
	GenerateSubsets(arr, k, index+1, subset, i)
}

func main() {
	arr := []int{1, 2, 3}
	k := 4

	subsets := GenerateSubsetsWithProductLessThanOrEqual(arr, k)

	fmt.Println(subsets)
}
