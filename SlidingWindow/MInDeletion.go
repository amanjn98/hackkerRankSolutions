package SlidingWindow

func MinimumDeletions(s string) int {
	end := 0
	start := 0
	count := 0
	for end < len(s) {
		if end-start+1 == 2 {
			if s[start:end+1] == "ba" {
				s = s[:start] + s[start+1:]
				count++
				start = start - 1
				end = end - 1
				continue
			} else {
				start++
			}
		}
		end++
	}
	return count
}
