package main

import "sync"

func pool(in chan int, numWorkers int, f func(int) int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for range numWorkers {
			wg.Add(1)
			go worker(in, out, f, &wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func worker(in, out chan int, f func(int) int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		out <- f(v)
	}
}
