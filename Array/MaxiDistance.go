package Array

import "math"

//You are given m arrays, where each array is sorted in ascending order.
//
//You can pick up two integers from two different arrays (each array picks one) and calculate the distance. We define the distance between two integers a and b to be their absolute difference |a - b|.
//
//Return the maximum distance.

//Example 1:
//
//Input: arrays = [[1,2,3],[4,5],[1,2,3]]
//Output: 4
//Explanation: One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.
//Example 2:
//
//Input: arrays = [[1],[1]]
//Output: 0

func maxDistance(arrays [][]int) int {
	minIndex, maxIndex := 0, 0
	min1, max1 := arrays[0][0], arrays[0][1]
	for i := 0; i < len(arrays); i++ {
		if max1 < arrays[i][len(arrays[i])-1] {
			max1 = arrays[i][len(arrays[i])-1]
			maxIndex = i
		}

		if min1 > arrays[i][0] {
			min1 = arrays[i][0]
			minIndex = i
		}
	}
	if minIndex == maxIndex {
		min2, max2 := max1, min1
		for i := 0; i < len(arrays); i++ {
			if min2 > arrays[i][0] && i != maxIndex {
				min2 = arrays[i][0]
			}

			if max2 < arrays[i][len(arrays[i])-1] && i != minIndex {
				max2 = arrays[i][len(arrays[i])-1]
			}
		}

		if int(math.Abs(float64(max1-min2))) < int(math.Abs(float64(max2-min1))) {
			return int(math.Abs(float64(max2 - min1)))
		}
		return int(math.Abs(float64(max1 - min2)))
	}
	return int(math.Abs(float64(max1 - min1)))
}
