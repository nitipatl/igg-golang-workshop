package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(*&namePointer)
	// printPointer(namePointer)
	fmt.Println(&name)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
