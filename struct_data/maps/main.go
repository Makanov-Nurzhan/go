package main

import "fmt"

func main() {
	product := map[string]int{}
	product["a"] = 1
	product["b"] = 2
	product["c"] = 3
	product["a"] = 2
	fmt.Println(product)
}
