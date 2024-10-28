package main

import (
	"fmt"
	"os"
)

func writeInBalanceSheet(text string) {
	textData := fmt.Sprint(text)
	os.WriteFile("fileone.txt", []byte(textData), 0644)
}

func readBalanceSheet() {
	value, _ := os.ReadFile("fileone.txt")
	fmt.Println("Enter Right name to get: ", string(value))
}

func bank() {
	var name string

	fmt.Println("Enter Your Name to Check Its Valid: ")
	fmt.Scan(&name)

	switch name {
	case "golang":
		fmt.Println("Yea you that one person")
		writeInBalanceSheet("Yea you that one person")
	case "":
		fmt.Println("Are you making fun on me?")
		readBalanceSheet()
		bank()

	default:
		fmt.Println("fuck you", name)
		readBalanceSheet()
		bank()
	}
}

func main() {
	bank()
}
