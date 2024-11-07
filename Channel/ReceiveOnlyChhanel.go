package Channel

import "fmt"

func main() {
	ch := add(10, 20)
	result := <-ch
	//ch<-1000 error since it is a receive only chanel
	fmt.Println(result)
}

func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result
	}()
	return ch
}
