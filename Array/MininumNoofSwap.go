package Array

//func minSwaps(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	j := len(s) - 1
//	count := 0
//	open := 0
//	closed := 0
//	for i := 0; i < len(s); i++ {
//		if s[i] == '[' {
//			open++
//		} else {
//			closed++
//		}
//		if open < closed {
//			open++
//			closed--
//			j--
//			count++
//		}
//	}
//	return count
//}

func minSwaps(s string) int {
	open := 0

	for _, v := range s {
		if v == '[' {
			open++
		} else {
			if open > 0 {
				open--
			} else {
				open++
			}
		}
	}
	return (open + 1) / 2
}
