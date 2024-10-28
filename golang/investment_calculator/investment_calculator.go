package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to investment calculator")
	var investmentAmount int
	var expectedReturn float64
	var years int
	var userName string
	fmt.Print("Enter a your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter a investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter a expected Returnnt: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Enter Your spending Year: ")
	fmt.Scan(&years)

	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturn/100), float64(years))
	fmt.Println(userName+" Your futre value is ", futureValue)
}
