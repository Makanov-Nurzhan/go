package main

import (
	"context"
	"fmt"
	"time"
)

func sourceA(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Source A finished", ctx.Err())
			close(ch)
			return
		case ch <- "Source A":
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func sourceB(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Source B finished", ctx.Err())
			close(ch)
			return
		case ch <- "Source B":
			time.Sleep(time.Second * 1)
		}
	}
}

func fanIn(chA, chB <-chan string, mergedCh chan string) {
	defer close(mergedCh)
	for chA != nil || chB != nil {
		select {
		case msgA, ok := <-chA:
			if !ok {
				chA = nil
				continue
			}
			mergedCh <- msgA
		case msgB, ok := <-chB:
			if !ok {
				chB = nil
				continue
			}
			mergedCh <- msgB
		}

	}
}

func main() {
	chA := make(chan string)
	chB := make(chan string)
	mergedCh := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go sourceA(ctx, chA)
	go sourceB(ctx, chB)
	go fanIn(chA, chB, mergedCh)

	for v := range mergedCh {
		fmt.Println(v)
	}
	fmt.Println("All sources finished")
}
