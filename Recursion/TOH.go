// Tower of Hanoi
//Lord Bhrama instructed priest to move the disk(n=64) from source to destination pole using helper repo .
//Once that task is completed, he will start the universe cycle once again
package Recursion

import "fmt"

func TOH(s int, h int, d int, n int) {
	if n == 1 {
		fmt.Print("Move disk from")
		fmt.Print(s)
		fmt.Print(" to ")
		fmt.Print(d)
		fmt.Println()
		return
	}
	TOH(s, d, h, n-1) // Moving n-1 disk from source to helper pole
	fmt.Print("Move disk from")
	fmt.Print(s)
	fmt.Print(" to ")
	fmt.Print(d)
	fmt.Println()
	TOH(h, s, d, n-1)
}

func CountTOH(s int, h int, d int, n int, count *int) {
	*count = *count + 1
	if n == 1 {
		fmt.Print("Move disk from")
		fmt.Print(s)
		fmt.Print(" to ")
		fmt.Print(d)
		fmt.Println()
		return
	}
	CountTOH(s, d, h, n-1, count) // Moving n-1 disk from source to helper pole
	fmt.Print("Move disk from")
	fmt.Print(s)
	fmt.Print(" to ")
	fmt.Print(d)
	fmt.Println()
	CountTOH(h, s, d, n-1, count)
}
