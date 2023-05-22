package main

import (
	//"./Graph"
	"./Recursion"
)

func main() {
	//matrix := [][]rune{
	//	{'O', 'M', 'O', 'O', 'X'},
	//	{'O', 'X', 'X', 'O', 'M'},
	//	{'O', 'O', 'O', 'O', 'O'},
	//	{'O', 'X', 'X', 'X', 'O'},
	//	{'O', 'O', 'M', 'O', 'O'},
	//	{'O', 'X', 'X', 'M', 'O'},
	//}
	Recursion.PrintSubset("aab", "")
}

// Send and receive integers to a channel using go routines.
// In the main function call a function "generator" which will add 10 integers to a channel; return the channel from the functions.
// In the main function call a function "receiver"; pass the channel to this function and print all the 10 integers.
//func main() {
//	ch := make(chan int, 1)
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go generator(ch, wg)
//	wg.Add(1)
//	go reciever(ch, wg)
//	wg.Wait()
//	close(ch)
//}
//
//func generator(ch chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	add := 0
//	for i := 0; i < 10; i++ {
//		fmt.Println("Writing data to a channel")
//		add = add + i
//	}
//	ch <- add
//
//}
//
//func reciever(ch chan int, wg *sync.WaitGroup) {
//	for i := <-ch; i >= 0; i-- {
//		fmt.Println(i)
//	}
//	wg.Done()
//
//}
