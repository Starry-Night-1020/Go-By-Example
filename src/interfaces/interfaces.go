package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (rec rect) area() float64 {
	return rec.width * rec.height
}

func (rec rect) perim() float64 {
	return rec.width*2 + rec.height*2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * 2
}

func PrintValue() {
	var r rect = rect{width: 10, height: 5}
	fmt.Println(r)

	var c circle = circle{radius: 10}
	fmt.Println(c)

	measure(r)
	measure(c)
}