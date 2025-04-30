package main

import "fmt"

func GenerateWithChannels(start, end int) <-chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)
		for num := start; num <= end; num++ {
			outputCh <- num
		}
	}()

	return outputCh
}

func main() {
	for num := range GenerateWithChannels(10, 20) {
		fmt.Println(num)
	}
}
