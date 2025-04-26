package main

import (
	"fmt"
	"sync"
)

func split(channel <-chan string, n int) []<-chan string {
	outChannels := make([]chan string, n)
	for i := 0; i < n; i++ {
		outChannels[i] = make(chan string)
	}
	go func() {
		idx := 0
		for value := range channel {
			outChannels[idx] <- value
			idx = (idx + 1) % n
		}
		for _, channel := range outChannels {
			close(channel)
		}
	}()

	resultChannels := make([]<-chan string, n)
	for i := 0; i < n; i++ {
		resultChannels[i] = outChannels[i]
	}

	return resultChannels
}
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
	splittedChs := split(channel, n)
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			for data := range splittedChs[idx] {
				out <- fmt.Sprintf("send -> %s", data)
			}
		}(i)
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
