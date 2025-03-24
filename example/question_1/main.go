// Write a function maxOfTwo that takes two integers and returns the larger one.

package main

import "fmt"

func maxOfTwo() {

	var number1 int
	var number2 int
	var max int

	fmt.Print("Enter no 1: ")
	fmt.Scan(&number1)

	fmt.Print("Enter no 2: ")
	fmt.Scan(&number2)

	if number1 > number2 {
		max = number1
		fmt.Print("the bigger no is number1: ", max)
	} else {
		max = number2
		fmt.Print("the bigger no is number2: ", max)
	}

}

func main() {

	maxOfTwo()
}
