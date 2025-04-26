package main

import (
	"fmt"
	"sync"
)

func SplitChannel(channel <-chan int, n int) []<-chan int {
	outChannels := make([]chan int, n)
	for i := 0; i < n; i++ {
		outChannels[i] = make(chan int)
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

	resultChannels := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		resultChannels[i] = outChannels[i]
	}

	return resultChannels
}
func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	channels := SplitChannel(channel, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1:", value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2:", value)
		}
	}()

	wg.Wait()
}
