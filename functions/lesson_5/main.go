package main

import "fmt"

func main() {

	factorial := factorial(3)
	fmt.Println(factorial)

}

// func factorial(num int) int {

// 	fac := 1

// 	for i := 1; i <= num; i++ {
// 		fac = fac * i
// 	}

// 	return fac
// }

// recursion style
func factorial(num int) int {

	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
