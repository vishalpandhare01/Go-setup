/*
remove duplicate element from array
Input = [1,2,3,1,,5,2,3,2,5,1]
output = [3 2 5 1]
*/
package main

import "fmt"

func main() {
	// your code here
	// for i
	//for j = i + 1
	// isDuplicate = false
	// if !isDuplicate : output = []
	array := []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5}
	output := []int{}
	for i := 0; i < len(array); i++ {
		var isDuplicate = false
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			output = append(output, array[i])
		}
	}

	fmt.Println("output", output)
}
