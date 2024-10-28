package main

import (
	"fmt"
)

func main() {
	var arrary []int
	// arrary = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if arrary == nil {
		fmt.Println("Empty array if array == nil (in this case var arrary []int)", arrary)
		return
	}
	fmt.Println("array", arrary)
	fmt.Println("len", len(arrary))
	fmt.Println("cap", cap(arrary))
	arrary = arrary[1:]
	arrary = arrary[:1]
	fmt.Println("array", arrary)
	fmt.Println("len", len(arrary))
	fmt.Println("cap", cap(arrary))

}
