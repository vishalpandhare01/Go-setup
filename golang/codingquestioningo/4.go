package main

import "fmt"

func main() {
	array := []int{9, 6, 3, 7, 4, 1}
	fmt.Println("Input", array)
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	fmt.Println("Output", array)
}
