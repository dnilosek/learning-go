package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("I am a %T and my area is %f\n", s, s.area())
}

func main() {
	sq := square{
		sideLength: 1.0,
	}
	cir := circle{
		radius: 1.0,
	}

	info(sq)
	info(cir)
}
