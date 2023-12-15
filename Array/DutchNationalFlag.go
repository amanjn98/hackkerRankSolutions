package Array

//Given an array A[] consisting of only 0s, 1s, and 2s.
//The task is to write a function that sorts the given array.
//The functions should put all 0s first, then all 1s and all 2s in last.

//Input: {0, 1, 2, 0, 1, 2}
//Output: {0, 0, 1, 1, 2, 2}

func sort012(arr []int) {
	l := 0
	r := len(arr) - 1
	i := 0
	for i < len(arr)-1 && i <= r {
		if arr[i] == 0 {
			arr[i], arr[l] = swap(arr[i], arr[l])
			i++
			l++
		} else if arr[i] == 2 {
			arr[i], arr[r] = swap(arr[i], arr[r])
			r--
		} else {
			i++
		}
	}
}

func swap(a int, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}
