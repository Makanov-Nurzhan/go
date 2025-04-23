package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r := Rectangle{10, 10}
	c := Circle{10}
	shapes := []Shape{r, c}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
