package main

import "fmt"

func main() {
	var array = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array", array)
	NewArray := append(array, 10)
	fmt.Println("Array", array)
	fmt.Println("NewArray", NewArray)
}
