package Dp

import "fmt"

//You are given an array of integers stones where stones[i] is the weight of the ith stone.
//
//We are playing a game with the stones. On each turn, we choose any two stones and smash them together. Suppose the stones have weights x and y with x <= y. The result of this smash is:
//
//If x == y, both stones are destroyed, and
//If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
//At the end of the game, there is at most one stone left.
//
//Return the smallest possible weight of the left stone. If there are no stones left, return 0.

// It can be solved using the concept of minimum subset sum problem

func MinSubsetSum(weight []int) {
	sum := 0
	for i := 0; i < len(weight); i++ {
		sum = sum + weight[i]
	}
	dp := SubsetSumForMin(sum, weight)
	var arr []int
	for i := 0; i <= sum/2; i++ {
		if dp[len(weight)][i] != false {
			arr = append(arr, i)
		}
	}
	mn := 999999
	for i := 0; i < len(arr); i++ {
		r := (sum - 2*arr[i])
		if mn > r {
			mn = r
		}
	}
	fmt.Println(mn)
}

func SubsetSumForMin(sum int, arr []int) [][]bool {
	dp := make([][]bool, len(arr)+1)
	for i := 0; i <= len(arr); i++ {
		dp[i] = make([]bool, sum+1)
		for j := 0; j <= sum; j++ {
			if i == 0 {
				dp[i][j] = false
			}
			if j == 0 {
				dp[i][j] = true
			}
		}
	}
	for i := 1; i <= len(arr); i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp
}
