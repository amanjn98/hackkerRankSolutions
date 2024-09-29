package Array

import "fmt"

func generator(n int) {
	start := n * (n + 1) / 2
	fmt.Println(start)
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(start + j - i)
			fmt.Print(" ")
		}
		fmt.Println()
		start = start - (i)
	}
}
