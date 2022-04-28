package main

import (
	"fmt"
)

type triangle struct {
	height float32
	width  float32
}

type circle struct {
	radius float32
}

type form interface {
	area() float32
}

func escreverArea(f form) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func (t triangle) area() float32 {
	return t.width * t.height
}

func (c circle) area() float32 {
	return c.radius * 2
}

func main() {
	fmt.Println("Interfaces")

	t := triangle{10, 8}
	escreverArea(t)

	c := circle{20}
	escreverArea(c)
}
