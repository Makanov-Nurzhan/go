package main

import "fmt"

func Transform(channel <-chan int, action func(int) int) <-chan int {
	outChannel := make(chan int)
	go func() {
		defer close(outChannel)
		for v := range channel {
			outChannel <- action(v)
		}
	}()

	return outChannel
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 5; i++ {
			channel <- i
		}
	}()

	mul := func(value int) int {
		return value * value
	}
	for numbers := range Transform(channel, mul) {
		fmt.Println(numbers)
	}

}
