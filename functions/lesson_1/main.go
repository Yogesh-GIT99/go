package main

import "fmt"

type transfntype func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	doubleNumber := transformNumber(&numbers, double)
	tripleNumber := transformNumber(&numbers, triple)

	fmt.Println(doubleNumber)
	fmt.Println(tripleNumber)

}

func transformNumber(number *[]int, transform transfntype) []int {

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
