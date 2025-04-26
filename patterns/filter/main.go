package main

import "fmt"

func Transform(channel <-chan int, action func(int) bool) <-chan int {
	outChannel := make(chan int)
	go func() {
		defer close(outChannel)
		for v := range channel {
			if action(v) {
				outChannel <- v
			}
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

	isOdd := func(value int) bool {
		return value%2 == 0
	}
	for numbers := range Transform(channel, isOdd) {
		fmt.Println(numbers)
	}

}
