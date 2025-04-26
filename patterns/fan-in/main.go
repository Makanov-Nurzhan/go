package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	outChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, channel := range channels {
		go func() {
			defer wg.Done()
			for value := range channel {
				outChannel <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outChannel)
	}()

	return outChannel

}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go func() {
		defer func() {
			close(channel1)
			close(channel2)
			close(channel3)
		}()
		for i := 0; i < 10; i++ {
			channel1 <- i
			channel2 <- i + 1
			channel3 <- i + 2
		}
	}()

	for value := range mergeChannels(channel1, channel2, channel3) {
		fmt.Println(value)
	}

}
