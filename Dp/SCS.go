package Dp

import "strings"

//Shortest common subsequence

func PrintLCS(a string, b string, m int, n int) string {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				if strings.EqualFold(string(a[i-1]), string(b[j-1])) {
					dp[i][j] = 1 + dp[i-1][j-1]
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}
	i := m
	j := n
	s := ""
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			s = string(a[i-1]) + s
			i--
			j--
		} else {
			if dp[i][j-1] > dp[i-1][j] { // since we need to take all elements in supersequence
				s = string(b[j-1]) + s
				j = j - 1
			} else {
				s = string(a[i-1]) + s
				i = i - 1
			}
		}
	}
	// Add Remaining elements
	for i > 0 {
		s = string(a[i-1]) + s
		i--
	}
	for j > 0 {
		s = string(b[j-1]) + s
		j--
	}
	return s
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
