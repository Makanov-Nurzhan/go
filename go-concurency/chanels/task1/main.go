package main

import (
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()
	select {
	case <-ch:
	case <-time.After(3 * time.Second):
	}
}

func main() {
	predictableTimeWork()
}
