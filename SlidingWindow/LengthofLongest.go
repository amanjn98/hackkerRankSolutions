package SlidingWindow

//Given a string s, find the length of the longest substring without repeating characters.

func LengthOfLongestSubstring(s string) int {
	position := map[rune]int{}
	if len(s) <= 1 {
		return len(s)
	}
	i := 0
	j := 0
	max := 0
	length := 0
	for j < len(s) {
		_, ok := position[rune(s[j])]
		if !ok {
			position[rune(s[j])] = 1
		} else {
			position[rune(s[j])]++
		}
		for position[rune(s[j])] > 1 {
			position[rune(s[i])] = position[rune(s[i])] - 1
			i++
		}
		length = j - i + 1
		if length > max {
			max = length
		}
		j++
	}
	return max
}
