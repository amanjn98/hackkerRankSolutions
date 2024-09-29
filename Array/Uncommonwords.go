package Array

import "strings"

//A sentence is a string of single-space separated words where each word consists only of lowercase letters.
//
//A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
//
//Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.

//Example 1:
//
//Input: s1 = "this apple is sweet", s2 = "this apple is sour"
//
//Output: ["sweet","sour"]
//
//Explanation:
//
//The word "sweet" appears only in s1, while the word "sour" appears only in s2.
//
//Example 2:
//
//Input: s1 = "apple apple", s2 = "banana"
//
//Output: ["banana"]

func uncommonFromSentences(s1 string, s2 string) []string {
	var sb strings.Builder
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	s1Array := strings.Split(sb.String(), " ")
	wordCount := make(map[string]int)
	for i := 0; i < len(s1Array); i++ {
		wordCount[s1Array[i]]++
	}
	result := make([]string, 0)
	for key, word := range wordCount {
		if word == 1 {
			result = append(result, key)
		}
	}
	return result
}
