package Dp

//You are given an array of binary strings strs and two integers m and n.
//
//Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
//
//A set x is a subset of a set y if all elements of x are also elements of y.
//Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
//Output: 4
//Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
//Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
//{"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.

func FindMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = 0
		}
	}
	for _, str := range strs {
		zeros, ones := getCountOfZeros(str)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zeros][j-ones])
			}
		}
	}
	return dp[m][n]
}

func getCountOfZeros(strs string) (int, int) {
	zero, one := 0, 0
	for i := 0; i < len(strs); i++ {
		if strs[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}
