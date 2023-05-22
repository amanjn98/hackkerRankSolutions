package Recursion

func PrintUniqueSubset(in string, op string, ipArray *[]string) {
	if len(in) == 0 {
		*ipArray = append(*ipArray, op)
		return
	}
	PrintUniqueSubset(in[1:], op, ipArray)
	PrintUniqueSubset(in[1:], op+string(in[0]), ipArray)
}

func RemoveDuplicates(ipArray []string) []string {
	values := map[string]bool{}
	result := []string{}
	for i := 0; i < len(ipArray); i++ {
		_, ok := values[ipArray[i]]
		if !ok {
			values[ipArray[i]] = true
			result = append(result, ipArray[i])
		}

	}
	return result
}