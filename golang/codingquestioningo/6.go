/*
check the given number is prime or not
*/
package main

import "fmt"

func IsPrime(number int) bool {
	//if number divided by 1 and its self is prime number
	// we divide number 2 to less than its self if divisible any of one then its non prime
	//if not its prime
	if number == 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	//write your code here
	for i := 1; i < 10; i++ {
		fmt.Println(i, IsPrime(i))
	}
}
