package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("(){{[]"))
}

func isValid(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {

		switch ch {
		case '(', '[', '{':
			// push
			stack = append(stack, ch)

		case '}', ')', ']':
			// check and pop
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]

		default:
			return false

		}
	}
	return len(stack) == 0
}
