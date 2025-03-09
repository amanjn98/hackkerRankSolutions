//Longest Common Prefix
//Write a function to find the longest common prefix string amongst an array of strings.If there is no common prefix, return an empty string "". 
//Example 1:Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
    if len(strs)==1{
        return string(strs[0])
    }
    temp:=201
    for i:=0;i<len(strs)-1;i++{
        j:=0
		for j < len(strs[i]) && j < len(strs[i+1]) && strs[i][j] == strs[i+1][j] {
			j++
		}
       if temp>j{
        temp=j
       }
    }
    if temp==0{
        return ""
    }
    var tempstr string
    for i:=0;i<temp;i++{
     tempstr+=string(strs[0][i])
    }
    return tempstr
}
