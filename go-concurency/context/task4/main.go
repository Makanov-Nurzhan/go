package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sourceA(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg := fmt.Sprintf("A: %s", time.Now().Format("15:04:05"))
		select {
		case <-ctx.Done():
			fmt.Println("Source A finished", ctx.Err())
			return
		case ch <- msg:
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func sourceB(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg := fmt.Sprintf("B: %s", time.Now().Format("15:04:05"))
		select {
		case <-ctx.Done():
			fmt.Println("Source B finished", ctx.Err())
			return
		case ch <- msg:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	wg.Add(2)
	go sourceA(ctx, ch, &wg)
	go sourceB(ctx, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("All sources finished")
}
