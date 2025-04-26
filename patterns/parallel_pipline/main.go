package main

import (
	"fmt"
	"sync"
)

func parse(channle <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for line := range channle {
			out <- fmt.Sprintf("parsed -> %s", line)
		}
	}()

	return out
}

func send(channel <-chan string, n int) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for data := range channel {
				out <- fmt.Sprintf("send -> %s", data)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	channel := make(chan string)

	go func() {
		defer close(channel)
		for i := 0; i < 5; i++ {
			channel <- "value"
		}
	}()

	for value := range send(parse(channel), 2) {
		fmt.Println(value)
	}
}
