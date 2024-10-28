package main

import "fmt"

//swap the number without third variable

func main() {
	a := 10
	b := 20
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	// a = b - a
	// b = a
	// a = a + b

	b, a = a, b
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
