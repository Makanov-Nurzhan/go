package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	v := <-ch
	fmt.Println(v)
}
