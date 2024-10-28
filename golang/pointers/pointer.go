package main

import "fmt"

func main() {
	var name string
	name = "vishal"
	var pointer = &name

	fmt.Println("name: ", name)
	name = "raj"
	*pointer = "pushparaj"
	fmt.Println("pointer store address: ", pointer)
	fmt.Println("pointer store value : ", *pointer)
	fmt.Println("name: ", name)
}

//pointer store  variable address
//its make duplicate variale
//pointer variable cant redeclear we can declear by using * its beahve actual variable
