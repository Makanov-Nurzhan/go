package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomWait() int {
	workSeconds := rand.Intn(5 + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func main() {
	ch := make(chan int)
	totalWorkSeconds := 0
	start := time.Now()
	for range 100 {
		go func() {
			seconds := randomWait()
			ch <- seconds
		}()

	}

	for range 100 {
		totalWorkSeconds += <-ch
	}

	mainSeconds := time.Since(start)

	fmt.Println("main:", mainSeconds)
	fmt.Println("total:", totalWorkSeconds)
}
