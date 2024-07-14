package Stack

import (
	"fmt"
	"strings"
)

//You are given a string s and two integers x and y. You can perform two types of operations any number of times.
//
//Remove substring "ab" and gain x points.
//For example, when removing "ab" from "cabxbae" it becomes "cxbae".
//Remove substring "ba" and gain y points.
//For example, when removing "ba" from "cabxbae" it becomes "cabxe".
//Return the maximum points you can gain after applying the above operations on s.

//Example 1:
//
//Input: s = "cdbcbbaaabab", x = 4, y = 5
//Output: 19
//Explanation:
//- Remove the "ba" underlined in "cdbcbbaaabab". Now, s = "cdbcbbaaab" and 5 points are added to the score.
//- Remove the "ab" underlined in "cdbcbbaaab". Now, s = "cdbcbbaa" and 4 points are added to the score.
//- Remove the "ba" underlined in "cdbcbbaa". Now, s = "cdbcba" and 5 points are added to the score.
//- Remove the "ba" underlined in "cdbcba". Now, s = "cdbc" and 5 points are added to the score.
//Total score = 5 + 4 + 5 + 5 = 19.
//Example 2:
//
//Input: s = "aabbaaxybbaabb", x = 5, y = 4
//Output: 20

func MaximumGain(s string, x int, y int) int {
	substring1 := ""
	substring2 := ""
	found := true
	points := 0
	if x > y {
		substring1 = "ab"
		substring2 = "ba"
	} else {
		substring1 = "ba"
		substring2 = "ab"
	}
	for found != false {
		index := strings.Index(s, substring1)
		if index == -1 {
			found = false
		} else {
			s = s[0:index] + s[index+2:]
			fmt.Println(s)
			if x > y {
				points = points + x
			} else {
				points = points + y
			}
		}
	}
	found = true
	for found != false {
		index := strings.Index(s, substring2)
		if index == -1 {
			found = false
		} else {
			s = s[0:index-1] + s[index+1:]
			fmt.Println(s)
			if x > y {
				points = points + y
			} else {
				points = points + x
			}
		}
	}
	return points
}

func maximumGain(s string, x int, y int) int {
	ans, curr := 0, 0

	if x >= y {
		curr = 0
		curr, s = helper(s, "ab", x)
		ans += curr
		curr, _ = helper(s, "ba", y)
		ans += curr
	} else {
		curr = 0
		curr, s = helper(s, "ba", y)
		ans += curr
		curr, _ = helper(s, "ab", x)
		ans += curr
	}
	return ans
}

func helper(s string, sub string, x int) (int, string) {
	var score int
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == sub[1] && len(stack) > 0 && stack[len(stack)-1] == sub[0] {
			stack = stack[:len(stack)-1] // pop the last element
			score += x
		} else {
			stack = append(stack, s[i])
		}
	}
	return score, string(stack)
}
