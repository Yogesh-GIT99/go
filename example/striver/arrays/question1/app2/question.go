// Problem Statement: Given a matrix if an element in the matrix is 0 then you will have to set its entire column and row to 0 and then return the matrix..
// SET matrix zero
// better approach

package main

import "fmt"

var matrix [][]int = [][]int{{1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 0}}

// rows = {false, true, false, true }   col = {true, false, false, true}

func main() {

	fmt.Println(matrix)
	setZeroesBetter(matrix)
	fmt.Println(matrix)

}

func setZeroesBetter(matrix [][]int) {

	n, m := len(matrix), len(matrix[0])
	rows := make([]bool, n)
	col := make([]bool, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if matrix[i][j] == 0 {
				rows[i] = true
				col[j] = true
			}

		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if rows[i] || col[j] {
				matrix[i][j] = 0
			}

		}
	}

}
