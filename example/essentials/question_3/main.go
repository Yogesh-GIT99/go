package main

import "fmt"

func main() {

	fmt.Println(factorial(6))

}

func factorial(n int) int {

	var fac = 1

	for i := n; i > 0; i-- {
		// fmt.Print(i)
		fac = fac * i
	}

	return fac
}
