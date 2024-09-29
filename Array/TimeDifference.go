package Array

import (
	"strconv"
	"strings"
)

//Given a list of 24-hour clock time points in "HH:MM" format, return the minimum minutes difference between any two time-points in the list.
//
//
//Example 1:
//
//Input: timePoints = ["23:59","00:00"]
//Output: 1
//Example 2:
//
//Input: timePoints = ["00:00","23:59","00:00"]
//Output: 0

//Minimum Time Difference

func findMinDifference(timePoints []string) int {
	minutes := make([]bool, 24* 60)
	max:= 0
	min := 24*60
	for _, timePoint := range(timePoints){
		parts := strings.Split(timePoint, ":")
		h, _ := strconv.Atoi(parts[0])
		m, _ := strconv.Atoi(parts[1])
		temp := h*60 + m
		if temp > max{
			max = temp
		}
		if temp < min{
			min = temp
		}
		if minutes[temp]{
			return 0
		}
		minutes[temp] = true


	}
	result := min - max + 1440

	prev := min
	for i := min+1; i<= max; i++{
		if minutes[i]{
			diff := i - prev
			if diff < result {
				result = diff
			}
			prev = i
		}


	}
	return result
}
