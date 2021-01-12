package main

import "fmt"

type square struct {
	s float64
}

type triangle struct {
	h float64
	b float64
}

type shape interface {
	getArea() float64
}

func main() {
	sqr := square{s: 20.10}
	printArea(sqr)
	tri := triangle{}
	tri.h = 5.5
	tri.b = 1.3
	printArea(tri)

}

func (s square) getArea() float64 {
	return s.s * s.s
}

func (t triangle) getArea() float64 {
	return .5 * t.b * t.h
}

func printArea(s shape) {
	fmt.Printf("%f \n", s.getArea())
}
