package main

import (
	"context"
	"fmt"
	"time"
)

func orWithContext[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		out := make(chan T)
		go func() {
			defer close(out)
			select {
			case <-ctx.Done():
			}
		}()
		return out
	case 1:
		return channels[0]
	}

	mid := len(channels) / 2

	left := orWithContext(ctx, channels[:mid]...)
	right := orWithContext(ctx, channels[mid:]...)

	out := make(chan T)

	go func() {
		defer close(out)
		select {
		case <-ctx.Done(): // Отмена через внешний context
		case v := <-left:
			out <- v
		case v := <-right:
			out <- v
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	start := time.Now()

	select {
	case <-orWithContext(ctx,
		time.After(2*time.Hour),
		time.After(5*time.Minute),
		time.After(1*time.Second),
		time.After(1*time.Hour),
		time.After(10*time.Second),
	):
		fmt.Printf("Completed by channels after: %s\n", time.Since(start))
	case <-ctx.Done():
		fmt.Printf("Cancelled by context after: %s\n", time.Since(start))
	}
}
