package main

import "fmt"

func main() {

	fmt.Println(getNoZeroIntegers(40053))
}

func getNoZeroIntegers(n int) []int {
	hasZero := func(x int) bool {
		for x > 0 {
			if x%10 == 0 {
				return true
			}
			x /= 10
		}
		return false
	}

	for a := 1; a < n; a++ {
		b := n - a
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return nil // guaranteed not to happen per problem statement
}
