package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomWait() int {
	workSeconds := rand.Intn(5 + 1)

	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func main() {
	wg := &sync.WaitGroup{}
	locker := &sync.Mutex{}
	totalWorkSeconds := 0
	start := time.Now()
	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			seconds := randomWait()
			locker.Lock()
			totalWorkSeconds += seconds
			locker.Unlock()
		}()

	}
	wg.Wait()
	mainSeconds := time.Since(start)

	fmt.Println("main:", mainSeconds)
	fmt.Println("total:", totalWorkSeconds)
}
