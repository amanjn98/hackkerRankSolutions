package Array

//Given an array arr of n elements that is first strictly increasing and then maybe strictly decreasing,
//find the maximum element in the array.
//Note: If the array is increasing then just print the last element will be the maximum value.
//Input: array[]= {5, 10, 20, 15}
//Output: 20
//Explanation: The element 20 has neighbors 10 and 15, both of them are less than 20.

//Input: array[] = {10, 20, 15, 2, 23, 90, 67}
//Output: 20 or 90
//Explanation: The element 20 has neighbors 10 and 15, both of them are less than 20, similarly 90 has neighbors 23 and 67.

//The following corner cases give a better idea about the problem.
//
//If the input array is sorted in a strictly increasing order, the last element is always a peak element.
//For example, 50 is peak element in {10, 20, 30, 40, 50}.
//If the input array is sorted in a strictly decreasing order, the first element is always a peak element.
//100 is the peak element in {100, 80, 60, 50, 20}.
//If all elements of the input array are the same, every element is a peak element.

func getPeakElement(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	// Strictly decreasing
	if arr[0] > arr[1] {
		return 0
	}
	// Strictly increasing
	if arr[len(arr)-1] >= arr[len(arr)-2] {
		return len(arr) - 1
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] >= arr[i-1] && arr[i] >= arr[i+1] {
			return i
		}
	}
	return -1
}
