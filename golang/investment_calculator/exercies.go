package main

import "fmt"

func main() {
	fmt.Println("welcome to Profit calculator!!!")
	var revenue int
	var expenses float64
	var taxRate int
	fmt.Println("Enter your Revenue: ")
	fmt.Scan(&revenue)

	fmt.Println("Enter your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("Enter your Tax Rate: ")
	fmt.Scan(&taxRate)

	var profit = (expenses - float64(revenue) - float64(taxRate))
	fmt.Printf("hello %v\nvishal %v\n", "Your profit is: ", profit)

}
