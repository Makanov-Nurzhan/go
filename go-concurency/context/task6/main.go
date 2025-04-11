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
		select {
		case <-ctx.Done():
			fmt.Println("⛔ Source A завершён:", ctx.Err())
			return
		case ch <- fmt.Sprintf("🔴 A: %s", time.Now().Format("15:04:05")):
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func sourceB(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("⛔ Source B завершён:", ctx.Err())
			return
		case ch <- fmt.Sprintf("🔵 B: %s", time.Now().Format("15:04:05")):
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	chA := make(chan string)
	chB := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go sourceA(ctx, chA, &wg)
	go sourceB(ctx, chB, &wg)

	go func() {
		wg.Wait()
		close(chA)
		close(chB)
	}()

	for {
		select {
		case msg, ok := <-chA:
			if !ok {
				chA = nil
				continue
			}
			fmt.Println("📥", msg)

		case msg, ok := <-chB:
			if !ok {
				chB = nil
				continue
			}
			fmt.Println("📥", msg)

		case <-ctx.Done():
			fmt.Println("🛑 Таймаут — выходим из select")
			return
		}
	}
}
