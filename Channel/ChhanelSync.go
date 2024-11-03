package Channel

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	result := id * 2 // Perform some work

	ch <- result // Send the result to the channel
}

func main() {
	numWorkers := 3
	ch := make(chan int, numWorkers) // Creating a buffered channel
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch) // Close the channel when all workers have finished
	}()

	for result := range ch {
		fmt.Println(result) // Process the results
	}
}
