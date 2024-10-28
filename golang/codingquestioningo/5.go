/*
count the duplicate elements in array
input = ["raj","vishal","bisal","raj","bisal","raj",vishal"]
output= {"raj":3,"vishal":2,"bisal":2}

*/

package main

import "fmt"

func main() {
	//your code here
	//create object {}
	//for loop itrate
	//if el exist in obj
	//if exist count++ and store
	//else i will add with key name and count 1
	// so on.....

	var obj = make(map[string]int)
	var array = []string{"raj", "vishal", "bisal", "raj", "bisal", "raj", "vishal", "raj", "vishal", "bisal", "raj", "bisal", "raj", "vishal"}
	for i := 0; i < len(array); i++ {
		if (obj[array[i]] > 0) && (len(obj) > 0) {
			obj[array[i]]++
		} else {
			obj[array[i]] = 1
		}
	}

	fmt.Println(obj)
}
