package main

import (
	"fmt"
	"reflect"
	"time"
)

func or[T any](channels ...<-chan T) <-chan T {
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}

	doneCh := make(chan T)

	go func() {
		defer close(doneCh)

		cases := make([]reflect.SelectCase, len(channels))
		for i, ch := range channels {
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			}
		}

		// Ожидаем первое событие
		_, _, _ = reflect.Select(cases)
	}()

	return doneCh
}

func main() {
	start := time.Now()

	<-or(
		time.After(2*time.Hour),
		time.After(5*time.Minute),
		time.After(1*time.Second),
		time.After(1*time.Hour),
		time.After(10*time.Second),
	)

	fmt.Printf("Called after: %s\n", time.Since(start))
}
