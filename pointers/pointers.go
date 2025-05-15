package main

import "fmt"

func main() {

	age := 32

	agePointer := &age // To call address of pointer use &

	// var agePointer *int
	// agePointer = &age           ------> The above can also be defined like this ( type of pointer is *type, i.e for int pointers it is *int)

	fmt.Println("age: ", *agePointer) // To call a value using a pointer use *

	// adultage := getAdultAge(age) ------> old way

	// fmt.Println("adultAge: ", adultage)

	getAdultAge(agePointer)

	fmt.Println("adultAge: ", age)
}

// Old way:

// func getAdultAge(age int) int {

// 	return age - 18
// }

// Using pointer:

// func getAdultAge(age *int) int {

// 	return *age - 18
// }

// Using data mutation:

func getAdultAge(age *int) {

	*age = *age - 18
}
