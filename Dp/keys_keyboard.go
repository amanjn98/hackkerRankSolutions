package Dp

// 2 keys keyboard
//There is only one character 'A' on the screen of a notepad. You can perform one of two operations on this notepad for each step:
//Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
//Paste: You can paste the characters which are copied last time.
//Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.
//Example 1:
//
//Input: n = 3
//Output: 3
//Explanation: Initially, we have one character 'A'.
//In step 1, we use Copy All operation.
//In step 2, we use Paste operation to get 'AA'.
//In step 3, we use Paste operation to get 'AAA'.
//Example 2:
//
//Input: n = 1
//Output: 0

func MinSteps(n int) int {
	if n < 2 {
		return 0
	}
	result := 0
	i := 5
	for {
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
			result += 2
		} else if n%3 == 0 {
			n = n / 3
			result += 3
		} else if n%i == 0 {
			n = n / i
			result += i
		} else {
			i++
		}
	}
	return result
}

// NOtes : This is a simple solution and we can think of prime factorization problem
