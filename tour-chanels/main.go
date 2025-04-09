package main

import (
	"fmt"
	"mymodule/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}

		if v1 != v2 {
			return false
		}
	}
	return true

}

func main() {
	t1 := tree.New(2)
	t2 := tree.New(1)
	res := Same(t1, t2)
	fmt.Println(res)

}
