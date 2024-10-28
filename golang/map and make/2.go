package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Initial slice
	people := []Person{}

	// Adding a new element using append
	people = append(people, Person{"Charlie", 22})
	people = append(people, Person{"pharlie", 22})
	people = append(people, Person{"kharlie", 22})
	people = append(people, Person{"dharlie", 22})

	for i, v := range people {
		fmt.Println(i, v.Name, v.Age)
	}

	fmt.Println(people)
}
