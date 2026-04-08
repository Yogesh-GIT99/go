// Problem Statement: Given a matrix if an element in the matrix is 0 then you will have to set its entire column and row to 0 and then return the matrix..
// SET matrix zero
// brute force

package main

import "fmt"

var matrix [][]int = [][]int{{1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 0}}

func main() {

	fmt.Println(matrix)
	setZeroesBrute(matrix)
	fmt.Println(matrix)

}

func setZeroesBrute(matrix [][]int) {

	n, m := len(matrix), len(matrix[0])

	var temp [][]int
	temp = make([][]int, n)
	for i := range temp {
		temp[i] = make([]int, m)
		copy(temp[i], matrix[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if temp[i][j] == 0 {
				// make whole row 0
				for col := 0; col < m; col++ {
					matrix[i][col] = 0
				}

				// make whole column 0
				for row := 0; row < n; row++ {
					matrix[row][j] = 0
				}
			}
		}

	}

}
