package SlidingWindow

//Given a string you need to print the size of the longest possible substring that has exactly K unique characters.
//If there is no possible substring then print -1.

//Input:
//S = "aabacbebebe", K = 3
//Output:
//7
//Explanation:
//"cbebebe" is the longest substring with 3 distinct characters.

//Input:
//S = "aaaa", K = 2
//Output: -1
//Explanation:
//There's no substring with 2 distinct characters.

func LongestKSubstr(s string, k int) int {
	position := map[rune]int{}
	i := 0
	j := 0
	max := 0
	length := 0
	position[rune(s[0])] = 0
	for j < len(s) {
		if _, ok := position[rune(s[j])]; !ok {
			position[rune(s[j])]++
		}
		for len(position) > k {
			value, _ := position[rune(s[i])]
			position[rune(s[i])] = value - 1
			if position[rune(s[i])] == 0 {
				delete(position, rune(s[i]))
			}
			i++
		}
		position[rune(s[j])]++
		length = j - i + 1
		if length > max {
			max = length
		}
		j++
	}
	if len(s) == max {
		return -1
	}
	return max
}
