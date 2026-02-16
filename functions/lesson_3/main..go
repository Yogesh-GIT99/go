package main

import (
	"fmt"
)

type transfntype func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	doubleNumber := transformNumber(&numbers, func(numbers int) int { return numbers * 2 })

	fmt.Println(doubleNumber)

}

func transformNumber(number *[]int, transform transfntype) []int {

	double := []int{}

	for _, value := range *number {

		double = append(double, transform(value))
	}

	return double

}
