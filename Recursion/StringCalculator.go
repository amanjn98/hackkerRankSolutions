package Recursion

// Asked in Dunzo and ServiceNow

func PrintUniqueSubset(in string, op string, ipArray *[]string) {
	if len(in) == 0 {
		*ipArray = append(*ipArray, op)
		return
	}
	op1 := in[1:]
	*ipArray = append(*ipArray, op1)
	PrintUniqueSubset(in[1:], op1, ipArray)
	if len(in)-1 > 0 {
		op2 := in[:len(in)-1]
		*ipArray = append(*ipArray, op2)
		PrintUniqueSubset(in[:len(in)-1], op2, ipArray)
		op3 := in[1 : len(in)-1]
		*ipArray = append(*ipArray, op3)
		PrintUniqueSubset(in[1:len(in)-1], op3, ipArray)
	}
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

//Given a string s , a substring is defined as no empty string that can be obtained by any of the following conditions .
//
//1)Remove zero or more characters from the left side of the string
//2)Remove zero or more characters from the right side of the string
//3) Remove zero or more characters from the left side of the string and
//Remove zero or more characters from the right side of the string
//
//Generate a substring calculator using the above conditions

//def generate_substrings(s):
//substrings = []
//n = len(s)
//
//for i in range(n):
//for j in range(i + 1, n + 1):
//substrings.append(s[i:j])
//
//return substrings
//
//# Example usage:
//input_string = "abcdef"
//substrings = generate_substrings(input_string)
//for substring in substrings:
//print(substring)
