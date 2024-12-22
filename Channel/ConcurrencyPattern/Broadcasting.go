package ConcurrencyPattern

import "fmt"

func broadcaster(ch chan int, listeners []chan int) {
	for value := range ch {
		for _, listener := range listeners {
			listener <- value
		}
	}
	for _, listener := range listeners {
		close(listener)
	}
}

func BroadCasting() {
	source := make(chan int)
	listener1 := make(chan int)
	listener2 := make(chan int)

	go broadcaster(source, []chan int{listener1, listener2})

	go func() {
		for i := 1; i <= 5; i++ {
			source <- i
		}
		close(source)
	}()

	go func() {
		for v := range listener1 {
			fmt.Println("Listener 1 received:", v)
		}
	}()

	for v := range listener2 {
		fmt.Println("Listener 2 received:", v)
	}
}
