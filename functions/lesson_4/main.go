package main

import (
	"fmt"
)

type transfntype func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	doubleNumber := transformNumber(&numbers, func(numbers int) int { return numbers * 2 })

	double := creatingTransformers(2)
	triple := creatingTransformers(3)

	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(doubleNumber)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumber(number *[]int, transform transfntype) []int {

	double := []int{}

	for _, value := range *number {

		double = append(double, transform(value))
	}

	return double

}

func creatingTransformers(factor int) func(int) int { // factory function

	return func(number int) int { // creating anonymous function
		return number * factor
	}
}
