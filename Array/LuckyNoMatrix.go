package Array

import "fmt"

//Given an m x n matrix of distinct numbers, return all lucky numbers in the matrix in any order.
//
//A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.

//Example 1:
//
//Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
//Output: [15]
//Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column.
//Example 2:
//
//Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
//Output: [12]
//Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.
//Example 3:
//
//Input: matrix = [[7,8],[1,2]]
//Output: [7]
//Explanation: 7 is the only lucky number since it is the minimum in its row and the maximum in its column.

func luckyNumbers(matrix [][]int) []int {
	minArr := []int{}
	for i := 0; i < len(matrix); i++ {
		minNo := matrix[i][0]
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < minNo {
				minNo = matrix[i][j]
			}
		}
		minArr = append(minArr, minNo)
	}
	maxArr := []int{}
	for i := 0; i < len(matrix[0]); i++ {
		maxEle := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > maxEle {
				maxEle = matrix[j][i]
			}
		}
		maxArr = append(maxArr, maxEle)
	}
	fmt.Println(maxArr)
	fmt.Println(minArr)
	res := []int{}
	for i := 0; i < len(minArr); i++ {
		for j := 0; j < len(maxArr); j++ {
			if minArr[i] == maxArr[j] {
				res = append(res, minArr[i])
			}
		}
	}
	return res
}
