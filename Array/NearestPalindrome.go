package Array

import (
	"math"
	"strconv"
)

func NearestPalindromic(n string) string {
	// parse left part
	len := len(n)
	leftPart := 0
	if len%2 == 0 {
		leftPart, _ = strconv.Atoi(n[0 : len/2])
	} else {
		leftPart, _ = strconv.Atoi(n[0 : len/2+1])
	}
	nInt, _ := strconv.Atoi(n)

	// generate 5 palin
	palins := []int{}
	palins = append(palins, generatePalin(leftPart, len%2 == 0))   // 12321
	palins = append(palins, generatePalin(leftPart+1, len%2 == 0)) // 12421
	palins = append(palins, generatePalin(leftPart-1, len%2 == 0)) // 12221
	palins = append(palins, int(math.Pow(10, float64(len-1))-1))   // 9999
	palins = append(palins, int(math.Pow(10, float64(len))+1))     // 10001

	// compare diff with n
	minDiff := 0
	var ans int
	for _, palin := range palins {
		if palin == nInt {
			continue
		}
		diff := abs(palin - nInt)
		if minDiff == 0 || diff < minDiff {
			minDiff = diff
			ans = palin
		} else if diff == minDiff {
			if palin < ans {
				ans = palin
			}
		}
	}

	return strconv.Itoa(ans)
}

func generatePalin(leftPart int, even bool) int {
	palin := leftPart
	if !even {
		leftPart = leftPart / 10
	}
	for leftPart > 0 {
		palin = palin*10 + leftPart%10
		leftPart /= 10
	}
	return palin
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
