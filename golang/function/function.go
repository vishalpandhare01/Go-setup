package main

import "fmt"

func func2() {
	fmt.Println("hello world")
}

func func1(next func()) {
	next()
}

func main() {
	func1(func2)
}
