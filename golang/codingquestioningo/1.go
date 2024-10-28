// reverse string
// input: vishal
// output: lahsiv
package main

import (
	"fmt"
)

func main() {

	var input = "vishal"
	var output string
	for i := len(input); i > 0; i-- {
		output += string(input[i-1]) //(input[i-1]), it retrieves that character as a byte (or rune). thats why used string()
	}
	fmt.Println(output)

}

// v i s h a l
// 0 1 2 3 4 5

// 5 4 3 2 1 0
// l a h s i v
