package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("foo")
}

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("gorotine")
	}()
	wg.Add(1)
	go foo(wg)
	wg.Wait()
	fmt.Println("main")

}
