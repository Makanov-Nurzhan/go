package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := &sync.WaitGroup{}
	var money atomic.Int32
	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			money.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(money.Load())

}
