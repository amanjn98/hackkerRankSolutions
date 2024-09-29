package Array

func KthDistinct(arr []string, k int) string {
	mapCount := make(map[string]int, len(arr))
	for i := 0; i < len(arr); i++ {
		mapCount[arr[i]]++
	}
	count := 1
	for i := 0; i < len(arr); i++ {
		if mapCount[arr[i]] == 1 && k == count {
			return arr[i]
		} else if mapCount[arr[i]] == 1 {
			count++
		}
	}
	return ""
}
