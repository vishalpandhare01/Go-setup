package main

import "fmt"

func main() {
	variadicFunc(1, 2, 3, 4, 5, 6)
}

func variadicFunc(firstNumber int, numbers ...int) {
	fmt.Println("firstNumber", firstNumber)
	fmt.Println("array", numbers) //return number in array
}
