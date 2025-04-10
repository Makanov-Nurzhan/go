package main

import (
	"context"
	"fmt"
	"time"
)

func generator(ctx context.Context, ch chan int) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("⛔ Генератор остановлен:", ctx.Err())
			close(ch)
			return
		case ch <- i:
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	ch := make(chan int)
	go generator(ctx, ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("done")

}
