package Trial

import (
	"fmt"
	"sync"
)

// Send and receive integers to a channel using go routines.
// In the main function call a function "generator" which will add 10 integers to a channel; return the channel from the functions.
// In the main function call a function "receiver"; pass the channel to this function and print all the 10 integers.

func main() {
	var ch chan int
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go generator(ch, wg)
	wg.Wait()
	wg.Add(1)
	go reciever(ch, wg)
	wg.Wait()
}

func generator(ch chan int, i *sync.WaitGroup) {
	defer i.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func reciever(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		res, ok := <-ch
		if ok == false {
			fmt.Println("No value")
			break
		}
		fmt.Println(res)
	}
}
