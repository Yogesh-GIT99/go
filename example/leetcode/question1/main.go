package main

import "fmt"

func main() {

	fmt.Println(sumZero(7))

}

func sumZero(n int) []int {
	var e_output []int
	var o_output []int = []int{0}
	var output []int
	sum := 0

	if n%2 == 0 {
		output = e_output
	} else {
		output = o_output
	}

	for i := n; i > 1; i-- {
		j := -i
		output = append(output, i)
		output = append(output, j)

		if len(output) == n {
			for _, value := range output {
				sum += value
			}

			if sum == 0 {
				return output
			}
		}
	}
	return output
}
