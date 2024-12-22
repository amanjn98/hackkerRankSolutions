package ConcurrencyPattern

import "fmt"

// Data flows through a sequence of stages, where each stage performs an operation and passes the result to the next stage.
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			fmt.Printf("Actual no:%d \n", num)
			out <- num
		}
		close(out)
	}()
	return out
}

func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func Pipeline() {
	gen := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	multiply := multiply(gen)
	for num := range multiply {
		fmt.Printf("Result: %d\n", num)
	}
}
