package ConcurrencyPattern

import (
	"fmt"
)

// Fan-Out: Multiple goroutines are created to handle tasks concurrently.
// Fan-In: Results from multiple goroutines are combined into a single channel.
func Producer(jobs chan<- int, n int) {
	for i := 1; i <= n; i++ {
		jobs <- i
	}
	close(jobs)
}
func workerProcess(id int, jobs <-chan int, results chan<- int) {
	//defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * 2
	}
}
func FanOutFanIn() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	go Producer(jobs, 10)
	//wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		//wg.Add(1)
		go workerProcess(i, jobs, results)
	}
	// Fan-In: Collect results
	for i := 1; i <= 10; i++ {
		fmt.Println(<-results)
	}
}
