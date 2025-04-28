package main

import (
	"fmt"
	"time"
)

func or[T any](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	mid := len(channels) / 2

	left := or(channels[:mid]...)  // первая половина
	right := or(channels[mid:]...) // вторая половина

	out := make(chan T)

	go func() {
		defer close(out)
		select {
		case v := <-left:
			out <- v
		case v := <-right:
			out <- v
		}
	}()

	return out
}

func main() {
	start := time.Now()

	<-or(
		time.After(2*time.Hour),
		time.After(5*time.Minute),
		time.After(2*time.Second),
		time.After(1*time.Hour),
		time.After(10*time.Second),
	)

	fmt.Printf("Called after: %s\n", time.Since(start))
}
