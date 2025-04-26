package main

import "fmt"

func generate(values ...int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range values {
			out <- v
		}
	}()
	return out
}

func process(channel chan int, action func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range channel {
			out <- action(v)
		}
	}()
	return out
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	mul := func(value int) int {
		return value * value
	}
	for value := range process(generate(values...), mul) {
		fmt.Println(value)
	}
}
