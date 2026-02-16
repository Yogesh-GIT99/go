package main

import (
	"fmt"
)

type transfntype func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	numbers2 := []int{3, 6, 9}
	doubleNumber := transformNumber(&numbers, double)
	tripleNumber := transformNumber(&numbers, triple)

	fmt.Println(doubleNumber)
	fmt.Println(tripleNumber)

	trasformNumberfn1 := gettransformer(&numbers)
	transformNumberfn2 := gettransformer(&numbers2)

	newtransformNumbers1 := transformNumber(&numbers, trasformNumberfn1)
	newtransformNumbers2 := transformNumber(&numbers, transformNumberfn2)

	fmt.Println(newtransformNumbers1)
	fmt.Println(newtransformNumbers2)

}

func transformNumber(number *[]int, transform transfntype) []int {

	double := []int{}

	for _, value := range *number {

		double = append(double, transform(value))
	}

	return double

}

func gettransformer(number *[]int) transfntype { // functions returning values

	if (*number)[1] == 2 {
		return double
	} else {
		return triple
	}

}

func double(val int) int {
	return 2 * val
}

func triple(val int) int {
	return 3 * val
}
