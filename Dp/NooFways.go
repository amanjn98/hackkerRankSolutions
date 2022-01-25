package Dp

import "strings"

var dp [100][100]int
func NumberOfWays(corridor string) int {
	for i:=0;i<len(corridor);i++{
		for j:=0;j<len(corridor);j++{
			dp[i][j]=-1
		}
	}
	return Solve(corridor,0,len(corridor)-1)
}

func Solve(s string, i int, j int) int {
	c:=Count(s,i,j)
	if c< 2 {
			return 0
	}else if c==2{
		return 1
	}
	if i>=j{
		return 0
	}

	ans:=0
	right:=0
	for k:=1; k<=j-2;k++{
		left:=Count(s,i,k)
		if left==2 {
			right = Count(s, k+1, j)
		}
		if left==2 && right==2{
			ans=ans+left*right
		}

	}
	return ans

}

func Count(s string, i int, j int) int {
	if i>=j{
		return 0
	}
	count:=0
	c:=i
	 for l,sub:=range s{
		 if l>j || c>j{
			 break
		 }
		 if l==c && strings.EqualFold(string(sub),"S"){
			 c++
			 count++
		 }else if l==c && l<j{
			 c++
		 }


	}
	return count
}