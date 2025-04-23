package main

import (
	"runtime"
	"time"
)

type Counter struct {
	ticker          *time.Ticker
	goroutine_count int
	ticks           int
}

var tickTime = 100 * time.Millisecond

func NewCounter() *Counter {
	ticker := time.NewTicker(tickTime)

	counter := &Counter{
		ticker: ticker,
	}
	counter.count()
	return counter
}

func (c *Counter) count() {
	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.ticks++
				c.goroutine_count += runtime.NumGoroutine()
			}
		}
	}()
}

func (c *Counter) GetAverageGoroutineCount() int {
	if c.ticks == 0 {
		return 0
	}
	return c.goroutine_count / c.ticks
}
func (c *Counter) Stop() {
	c.ticker.Stop()
}
func (c *Counter) Restart() {
	c.goroutine_count = 0
	c.ticks = 0
	c.ticker.Reset(tickTime)
}
