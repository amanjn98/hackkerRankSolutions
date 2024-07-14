package Dp

//Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.
//
//In one step, you can delete exactly one character in either string.

//Example 1:
//
//Input: word1 = "sea", word2 = "eat"
//Output: 2
//Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
//Example 2:
//
//Input: word1 = "leetcode", word2 = "etco"
//Output: 4

func LCS(a string, b string) int {
	t := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		t[i] = make([]int, len(b))
		for j := 0; j <= len(b); j++ {
			if i == 0 || j == 0 {
				t[i][j] = 0
			} else {
				t[i][j] = -1
			}
		}
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = max(t[i][j-1], t[i-1][j])
			}
		}
	}
	return t[len(a)][len(b)]
}

func NoofSteps(a string, b string) int {
	del := len(a) - LCS(a, b)
	ins := len(b) - LCS(a, b)
	return del + ins
}
