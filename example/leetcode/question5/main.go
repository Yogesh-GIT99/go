package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isPalindrome(565))

}

func isPalindrome(x int) bool {

	stringA := strconv.Itoa(x)
	reverse := ""

	for _, word := range stringA {
		reverse = string(word) + reverse
	}

	return stringA == reverse
}
