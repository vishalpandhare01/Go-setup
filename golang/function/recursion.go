package main

import "fmt"

func printNumbers(number int) int {
	if number == 0 {
		return number
	} else {
		fmt.Println(number)
		return printNumbers(number - 1)
	}
}

func main() {
	fmt.Println(printNumbers(10))
}
