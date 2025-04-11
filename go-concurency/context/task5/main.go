package main

import (
	"context"
	"fmt"
	"time"
)

func sourceA(ctx context.Context, ch chan string) {
	for {
		msg := fmt.Sprintf("A: %s", time.Now().Format("15:04:05"))
		select {
		case <-ctx.Done():
			fmt.Println("Source A finished", ctx.Err())
			close(ch)
			return
		case ch <- msg:
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func sourceB(ctx context.Context, ch chan string) {
	for {
		msg := fmt.Sprintf("B: %s", time.Now().Format("15:04:05"))
		select {
		case <-ctx.Done():
			fmt.Println("Source B finished", ctx.Err())
			close(ch)
			return
		case ch <- msg:
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	chA := make(chan string)
	chB := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go sourceA(ctx, chA)
	go sourceB(ctx, chB)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Чтение завершено", ctx.Err())
			return
		case v, ok := <-chA:
			if !ok {
				chA = nil
				continue
			}
			fmt.Println(v)
		case v, ok := <-chB:
			if !ok {
				chB = nil
				continue
			}
			fmt.Println(v)
		}
	}

	fmt.Println("All sources finished")
}
