package main

import "fmt"

func main() {

	num := []int{6, 5, 12}

	sum := sumup(3, 5, 10)
	anothersum := sumup(num...) // slicing values into paramters

	fmt.Println(sum)
	fmt.Println(anothersum)
}

// regular way
// func sumup(numbers []int) int {

// 	sum := 0

// 	for _, value := range numbers {

// 		sum = sum + value
// 	}

// 	return sum

// }

func sumup(numbers ...int) int { // such parameters are called as variadic functions

	sum := 0

	for _, value := range numbers {

		sum = sum + value
	}

	return sum

}
