package main

import (
	"context"
	"fmt"
	"time"
)

func generate() chan int {
	in := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			in <- i
		}
		close(in)
	}()
	return in
}

var numWorkers = 10

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	now := time.Now()

	counter := NewCounter()
	for v := range fanin(ctx, fanout(generate(), numWorkers, square)) {
		fmt.Println(v)
	}
	counter.Stop()
	avgFaninGoroutines := counter.GetAverageGoroutineCount()

	timeFanin := time.Since(now)

	now = time.Now()
	counter.Restart()
	for v := range pool(generate(), numWorkers, square) {
		fmt.Println(v)
	}
	counter.Stop()
	avgPoolGoroutines := counter.GetAverageGoroutineCount()
	timePool := time.Since(now)
	fmt.Println("time fanin:", timeFanin)
	fmt.Println("time pool:", timePool)
	fmt.Println("avg fanin goroutines:", avgFaninGoroutines)
	fmt.Println("avg pool goroutines:", avgPoolGoroutines)
}
