package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type contextKey string

const userIDKey contextKey = "userID"

func processTask(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	//val := ctx.Value(userIDKey)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершён: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Воркер %d работает: \n", id)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, userIDKey, 123)
	wg := sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		go processTask(ctx, i+1, &wg)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("🛑 Отмена всех воркеров!")
	cancel()
	wg.Wait()
	fmt.Println("✅ Все воркеры завершены. Main завершён.")
}
