// Problem Statement: Given a matrix if an element in the matrix is 0 then you will have to set its entire column and row to 0 and then return the matrix..
// SET matrix zero
// optimal approach
package main

import "fmt"

var matrix [][]int = [][]int{{1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 0}}

// rows = {false, true, false, true }   col = {true, false, false, true}
var col0 int = 1

func main() {

	fmt.Println(matrix)
	setZeroesOptimal(matrix)
	fmt.Println(matrix)

}

func setZeroesOptimal(matrix [][]int) {
	rows, col := len(matrix), len(matrix[0])

	// mark rows and columns
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// update matrix
	for i := 1; i < rows; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}

		}
	}

	// handle first row
	if matrix[0][0] == 0 {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}

	// handle first column
	if col0 == 0 {
		for j := 0; j < rows; j++ {
			matrix[j][0] = 0
		}
	}

}
