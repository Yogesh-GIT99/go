package main

import "fmt"

// function as types
type transfntype func(int) int // defining function as type/alias to make parameter readable

// type anotherType func(int, string, map[string]int) (string, float64)      ----> this is where defining func as type becomes helpful.

func main() {

	numbers := []int{1, 2, 3, 4}
	doubleNumber := transformNumber(&numbers, double) // passing functions as parameter values
	tripleNumber := transformNumber(&numbers, triple)

	fmt.Println(doubleNumber)
	fmt.Println(tripleNumber)

}

// creating a generic function instead of creating two functions for double and triple.
func transformNumber(number *[]int, transform transfntype) []int { // defining general function parameter value

	double := []int{}

	for _, value := range *number {

		double = append(double, transform(value))
	}

	return double

}

func double(val int) int {
	return 2 * val
}

func triple(val int) int {
	return 3 * val
}
