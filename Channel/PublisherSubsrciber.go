package Channel

import (
	"fmt"
	"sync"
)

func squareWorker(id int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range input {
		square := num * num
		output <- square
	}
}

func main() {
	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go squareWorker(1, input, output, &wg)
	go squareWorker(2, input, output, &wg)

	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	go func() {
		defer close(output)
		wg.Wait()
	}()

	for square := range output {
		fmt.Println(square)
	}

	fmt.Println("All values processed")
}
