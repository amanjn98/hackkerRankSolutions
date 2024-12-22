package Channel

import (
	"fmt"
	"sync"
)

func printOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		ch <- i // Send odd number to channel
		fmt.Println("Odd:", i)
	}
	close(ch) // Close channel after sending all odd numbers
}

func printEven(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		_, ok := <-ch // Receive odd number from channel
		if ok {
			fmt.Println("Even:", i)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // Create an unbuffered channel

	wg.Add(2)

	go printOdd(ch, &wg)
	go printEven(ch, &wg)

	wg.Wait()
}
