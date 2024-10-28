package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SeparateNumbers(s string) {
	var num int
	var numStr string
	var num1 int
	num1 = 0
	if len(s) == 1 {
		fmt.Print("NO")
		return
	}
	for i := 0; i < len(s)/2; i++ {
		num = getIntFromString(s, i)
		numStr = strconv.Itoa(num)
		num1 = num
		for len(numStr) < len(s) {
			num1 = num1 + 1
			numStr = numStr + strconv.Itoa(num1)
		}
		if strings.EqualFold(s, numStr) {
			numStr = strconv.Itoa(num)
			break
		}
	}
	if len(numStr) < len(s) {
		fmt.Print("YES ")
		fmt.Print(numStr)
	} else {
		fmt.Print("NO")
	}
}

func getIntFromString(s string, length int) int {
	num := 0
	for i := 0; i <= length; i++ {
		newNum, _ := strconv.Atoi(string(s[i]))
		num = newNum + num*10
	}
	return num
}
