package main

import "fmt"

func main() {

	fmt.Println(sumUpTo(2))

}

func sumUpTo(n int) int {

	var sum = (n * (n + 1)) / 2
	return sum
}
