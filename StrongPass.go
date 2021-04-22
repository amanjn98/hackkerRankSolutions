package main

import "fmt"

// Complete the minimumNumber function below.
func minimumNumber(n int32, passw string) int32 {
	  // Return the minimum number of characters to make the passw strong

	requiredDigits:=0
    var digit,lower,special,upper bool

    for i:=0;i<len(passw);i++{
      if passw[i]>='A' && passw[i] <='Z'{
            upper=true
	  }else if passw[i]>='a' && passw[i] <='z'{
            lower=true
	  }else if passw[i]>='0' && passw[i] <='9'{
           digit=true
	  }else{
         special=true
	  }
	}

	if !digit{
		requiredDigits++
	}
	if !lower{
		requiredDigits++
	}
	if !upper{
		requiredDigits++
	}
	if !special{
		requiredDigits++
	}
	if requiredDigits+len(passw)<6{
		requiredDigits=6-len(passw)
	}
    return int32(requiredDigits)

}

func main() {
	a:="2bbbb"
	fmt.Println(minimumNumber(int32(len(a)),a))
}



