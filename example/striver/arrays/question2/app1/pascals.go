package main

import "fmt"

func main() {
	ele, row := printPascalsTraingle(5, 5, 3)
	fmt.Println("Element at position (r, c):", ele)
	fmt.Println("N-th row of Pascal’s triangle: ", row)

}

var ele int
var row []int

func printPascalsTraingle(N, r, c int) (int, []int) {

	fmt.Printf("First %d rows of Pascal’s triangle:\n", N)

	for i := 0; i < N; i++ {
		val := 1
		row = make([]int, N)
		for j := 0; j < i+1; j++ {

			fmt.Print(val)

			if r-1 == i && c-1 == j {
				ele = val
			}

			if i == N-1 {
				row[j] = val
			}

			val = val * (i - j) / (j + 1)

		}
		fmt.Println()
	}

	return ele, row
}
