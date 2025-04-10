package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sensor(ctx context.Context, wg *sync.WaitGroup, ch chan string, sensorId int) {
	defer wg.Done()
	for {
		msg := fmt.Sprintf("📡 Сенсор %d: %s", sensorId, time.Now().Format("15:04:05"))
		select {
		case <-ctx.Done():
			fmt.Printf("⛔ Сенсор %d завершён: %v\n", sensorId, ctx.Err())
			return
		case ch <- msg:
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go sensor(ctx, &wg, ch, i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("✅ Main завершён")
}
