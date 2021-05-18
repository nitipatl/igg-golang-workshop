package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// taey := person{
	// 	"Nitipat",
	// 	"L",
	// }
	// var taey person
	// taey.firstName = "Nitipat"
	// taey.lastName = "L"

	taey := person{
		firstName: "Nitipat",
		lastName:  "L",
		contact: contactInfo{
			email: "iamsvz@gmail.com",
			zip:   50200,
		},
	}

	fmt.Printf("%+v", taey)
}
