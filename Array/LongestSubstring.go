// Print Substring without any repeating character
package Array

import "strings"

//GEEKSFORGEEKS"Longest substring: "EKSFORG", "KSFORGE"

func longestSubstring(s string) string {
	s = strings.ToLower(s)
	sp := 0
	start := 0
	length := 0
	max := 0
	position := make([]int, 125)
	for i := 0; i < len(position); i++ {
		position[i] = -1
	}
	// starting point of current index
	position[s[0]] = 0
	for i := 1; i < len(s); i++ {
		if position[s[i]] == -1 {
			position[s[i]] = i
		} else {
			if position[s[i]] >= sp {
				length = i - sp
				if max < length {
					max = length
					start = sp
				}
				sp = position[s[i]] + 1
			}
			position[s[i]] = i
		}
		if max < i-sp {
			max = i - sp
			start = sp
		}
	}
	result := ""
	for i := start; i < start+max; i++ {
		result = result + string(s[i])
	}
	return result
}
