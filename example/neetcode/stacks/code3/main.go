package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	cal []int
}

func (s *Stack) push(val int) {
	s.cal = append(s.cal, val)
}

func (s *Stack) pop() int {
	val := s.cal[len(s.cal)-1]
	s.cal = s.cal[:len(s.cal)-1]
	return val
}

func evalRPN(tokens []string) int {

	stack := Stack{}
	result := 0

	for _, char := range tokens {

		if num, err := strconv.Atoi(char); err == nil {
			stack.cal = append(stack.cal, num)
			continue
		}

		a := stack.pop()
		b := stack.pop()

		switch char {

		case "+":
			result = b + a
		case "-":
			result = b - a
		case "*":
			result = b * a
		case "/":
			result = b / a

		}

		stack.cal = append(stack.cal, result)
	}
	return stack.pop()
}

func main() {

	var tokens = []string{"2", "3", "+", "1", "-"}
	fmt.Println(evalRPN(tokens))
}
