package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func startSource(ctx context.Context, id int, ch chan string) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("⛔ Source %d завершён: %v\n", id, ctx.Err())
			return
		case ch <- fmt.Sprintf("📡 Source %d: %s", id, time.Now().Format("15:04:05")):
			time.Sleep(time.Millisecond * time.Duration(300+rand.Intn(400)))
		}
	}
}

func fanInMany(channels []<-chan string, merged chan string) {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan string) {
			defer wg.Done()
			for v := range ch {
				merged <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(merged)
	}()

}

func main() {
	merged := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	numSources := 5
	sources := make([]chan string, numSources)
	for i := 0; i < numSources; i++ {
		sources[i] = make(chan string)
		go startSource(ctx, i+1, sources[i])
	}

	readableSources := make([]<-chan string, numSources)
	for i, ch := range sources {
		readableSources[i] = ch
	}

	go fanInMany(readableSources, merged)

	for v := range merged {
		fmt.Println(v)
	}

	fmt.Println("All sources finished")
}
