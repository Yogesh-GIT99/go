// You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

// The input string s is valid if and only if:

// Every open bracket is closed by the same type of close bracket.
// Open brackets are closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
// Return true if s is a valid string, and false otherwise.
package main

import "fmt"

func main() {
	fmt.Println(isValid("{{()}}"))
	fmt.Println(isValid("{{()}"))
	fmt.Println(isValid("{{())}}"))
}

func isValid(s string) bool {

	stack := []rune{}
	pair := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, ch := range s {

		// insert all opening brackets to slack
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
			continue
		}

		// Closing brackets but there are none
		if len(stack) == 0 {
			return false
		}

		// checking all opening brackets have closing brackets or return false
		if stack[len(stack)-1] != pair[ch] {
			return false
		}

		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
