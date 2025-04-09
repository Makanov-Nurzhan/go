package main

import (
	"fmt"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) chan int {
	ch2x := make(chan int)
	go func() {
		for i := range ch {
			time.Sleep(500 * time.Millisecond)
			ch2x <- i * 2
		}
		close(ch2x)
	}()

	return ch2x
}

func reader(ch chan int) {
	for v := range ch {
		fmt.Println("v= ", v)
	}
}
func main() {
	reader(doubler(writer()))
}
