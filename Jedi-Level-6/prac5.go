package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) calculate() float64 {
	return s.length * 2
}

func (c circle) calculate() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	calculate() float64
}

func info(s shape) {
	fmt.Println(s.calculate())
}

func main() {
	s := square{20}
	info(s)
	c := circle{20}
	info(c)
}
