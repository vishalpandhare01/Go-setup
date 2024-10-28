/* unpack values using (...) dots*/

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	newElement := []int{7, 8, 9, 0}
	array = append(array, newElement...)
	fmt.Println("array", array)
}
