package main

import "fmt"

func Average(slice []int) float64 {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return float64(sum) / float64(len(slice))
}

func main() {
	slice := []int{10, 20, 30, 40}
	fmt.Println(Average(slice))
}
