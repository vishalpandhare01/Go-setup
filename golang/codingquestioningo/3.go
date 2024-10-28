// reverse number
// input 12345
// output 54321
package main

import "fmt"

func main() {
	var input = 12345
	var output = 0
	for input > 0 {
		var n = input % 10
		output = (10 * output) + n
		fmt.Println(output)
		input = input / 10
	}

	fmt.Println("output", output)

}

//5/10
