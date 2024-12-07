package Channel

import (
	"fmt"
	"sync"
)

func producer(start int, end int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < end; i++ {
		ch <- i
	}
}
func Consumer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func CallProducerAndConsumer() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(0, 4, ch, &wg)
	go producer(5, 10, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()
	Consumer(ch)
}
