package main

import "fmt"

func findPossible(numcourses int,prequisite [][]int) bool{

	require:=make([]int,numcourses)
	learned:=make([]bool,numcourses)
	s:=make([][]int,numcourses)
	for _,v:=range prequisite{
		require[v[0]]++
		s[v[1]]=append(s[v[1]],v[0])
	}
	for {
		learn:=false
		for  i,v:=range require{
			if v<=0 && learned[i]==false {
				learn = true
				learned[i] = true

				for _, v := range s {
					require[v[1]]--
				}
			}
		}
		if learn==false{
			return false
		}
		return true
	}
}

func main() {
	p:=[][]int{{1,0}}
	fmt.Println(findPossible(2,p))
}