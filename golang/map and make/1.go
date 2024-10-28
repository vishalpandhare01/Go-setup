package main

import "fmt"

func main() {
	obj := map[string]string{
		"google": "https//:google.com",
		"Aws":    "https//:amazon.com",
	}
	obj["linkedin"] = "htttps//:linkedin.com"
	fmt.Println(obj)
	delete(obj, "Aws")
	fmt.Println(obj)

	//by creating make we can use array[0] to add element in array
	array := make([]string, 2, 5)
	array[0] = "ohh my god!"
	array = append(array, "check", "ok", "bhai")

	fmt.Println("array", array)

}
