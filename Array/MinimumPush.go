package Array

import "sort"

func minimumPushes(word string) int {
	mapCount := make([]int, 26)
	for i := 0; i < len(word); i++ {
		mapCount[int(word[i]-'a')]++
	}
	sort.Slice(mapCount, func(i, j int) bool {
		return mapCount[i] > mapCount[j]
	})
	pushes := 0
	for i := 0; i < len(mapCount); i++ {
		pushes += (i/8 + 1) * mapCount[i]
	}
	return pushes
}
