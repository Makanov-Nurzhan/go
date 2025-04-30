package main

import (
	"context"
	"fmt"
	"time"
)

type Promise[T any] struct {
	resultCh chan T
}
type Future[T any] struct {
	resultCh <-chan T
}

func NewPromise[T any]() *Promise[T] {
	return &Promise[T]{resultCh: make(chan T)}
}
func NewFuture[T any](resultCh <-chan T) Future[T] {
	return Future[T]{
		resultCh: resultCh,
	}
}

func (p *Promise[T]) Set(value T) {
	p.resultCh <- value
}

func (f *Future[T]) Get() T {

	return <-f.resultCh
}

func (f *Future[T]) GetWithTimeout() (T, error) {
	select {
	case value := <-f.resultCh:
		return value, nil
	case <-time.After(time.Second * 2):
		var value T
		return value, fmt.Errorf("timeout")
	}
}

func (f *Future[T]) GetWithContext(ctx context.Context) (T, error) {
	select {
	case value := <-f.resultCh:
		return value, nil
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	}
}

func (p *Promise[T]) GetFuture() Future[T] {
	return NewFuture(p.resultCh)
}

func main() {
	promise := NewPromise[string]()

	go func() {
		time.Sleep(time.Second)
		promise.Set("agreement")
	}()

	future := promise.GetFuture()
	value, err := future.GetWithTimeout()
	fmt.Println(value, err)
}
