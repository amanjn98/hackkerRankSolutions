package main

import (
	"fmt"
	"math"
)

func convertFive(n int) int {
	if n==0{
		return 5
	}
	num:=0
	count:=0
	for n!=0{
		i :=n%10
		if i==0{
			i=5
		}
		num=num+i*int(math.Pow(10, float64(count)))
		count++
		n=n/10
	}
	return num
}
//1552=1002+5*10^2+5*10^1
//decimalplace as 0 place
func convertFiveusingDecimalplaces(n int) int {
  if n==0{
	 return 5
  }
  result:=0
  decimalplace:=1
  number:=n
  for number>0{
	  if number%10==0{
		  result=result+5*decimalplace
	  }
	  decimalplace=decimalplace*10
	  number=number/10
  }
	return n+result
}

func main() {
	fmt.Println(convertFiveusingDecimalplaces(104))
}