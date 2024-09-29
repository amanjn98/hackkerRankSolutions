package Array

func mergeSort(items []int) []int {

	if len(items) < 2 {

		return items

	}

	first := mergeSort(items[:len(items)/2])

	second := mergeSort(items[len(items)/2:])

	return merge(first, second)

}

func merge(arr1 []int, arr2 []int) []int {
	final1 := []int{}
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			final1 = append(final1, arr1[i])
			i++
		} else {
			final1 = append(final1, arr2[j])
			j++
		}
	}
	if i < len(arr1) {
		final1 = append(final1, arr1[i:]...)
	}
	if j < len(arr2) {
		final1 = append(final1, arr2[j:]...)
	}
	return final1
}
