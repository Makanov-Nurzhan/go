package main

import "time"

// cpu intensive
func timeConsuming1() {
	counter := 0
	for i := 0; i < 10000000; i++ {
		counter++
	}
}

// I/O
func timeConsuming2() {
	time.Sleep(100 * time.Millisecond)
}

func timeConsuming() {
	timeConsuming1()
}
