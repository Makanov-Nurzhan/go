package main

import (
	"context"
	"sync"
)

func fanin(ctx context.Context, chans []chan int) chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func fanout(in chan int, numChans int, f func(int) int) []chan int {
	chans := make([]chan int, numChans)

	for i := range numChans {
		chans[i] = pipeline(in, f)
	}
	return chans

}

func pipeline(in chan int, f func(int) int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- f(v)
		}
		close(out)
	}()
	return out

}
