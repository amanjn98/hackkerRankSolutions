package ConcurrencyPattern

import (
	"fmt"
	"sync"
)

// A fixed number of worker goroutines process tasks from a shared queue.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
}

func WorkerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	wg := sync.WaitGroup{}
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
