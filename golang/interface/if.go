package main

import (
	"fmt"
	"math"
)

// Abstraction: You want to define a contract that multiple types can implement. This allows different types to be treated uniformly based on the methods they provide.
type Shape interface {
	Area() float64
}

// Polymorphism: You want to write functions that can accept different types that implement the same interface. This is useful for writing flexible and reusable code.

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

// Decoupling: You want to reduce dependencies between components. Interfaces allow you to change implementations without affecting the code that uses them.

// Key Differences:
// Structs are about data representation, while interfaces are about behavior.
// A struct can have fields and methods, while an interface only defines method signatures.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func main() {
	shapes := []Shape{Circle{Radius: 5}, Square{Side: 4}}
	for _, shape := range shapes {
		PrintArea(shape)
	}
}
