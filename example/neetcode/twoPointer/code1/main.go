package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(isPanlindrome("Was it a car or a cat I saw?"))

}

func isPanlindrome(s string) bool {

	s = strings.ToLower(s)

	i := 0
	j := len(s) - 1

	for i < j {

		for i < j && !isalnum(s[i]) {
			i++
		}

		for i < j && !isalnum(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isalnum(c byte) bool {

	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
