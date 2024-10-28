package main

import (
	"fmt"
)

type UserData struct {
	firstName string
	lastName  string
	dob       string
}

func main() {
	firstName := addUserDeteil("Enter Your first name: ")
	lastName := addUserDeteil("Enter Your last name: ")
	dob := addUserDeteil("Enter Your dob (dd/mm/yyyy): ")

	var data UserData

	data = UserData{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
	}

	// here is difrent way
	data.OutputUserData()

}

func (u *UserData) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.dob)
}

func addUserDeteil(text string) string {
	var val string
	fmt.Println(text)
	fmt.Scan(&val)
	return val
}
