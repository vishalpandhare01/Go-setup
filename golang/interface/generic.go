package main

import "fmt"

func main() {
	result := add(10, 20.5)
	fmt.Println(result)
}

func add[T int | float64](a, b T) T {
	return a + b
}
