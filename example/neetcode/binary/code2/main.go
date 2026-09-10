// You are given an m x n 2-D integer array matrix and an integer target.
// Each row in matrix is sorted in non-decreasing order.
// The first integer of every row is greater than the last integer of the previous row.
// Return true if target exists within matrix or false otherwise.
// Can you write a solution that runs in O(log(m * n)) time?

package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}

	fmt.Println(searchMatrix(matrix, 15))
}

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix)
	col := len(matrix[0])

	low := 0
	high := row*col - 1

	for low <= high {

		mid := low + (high-low)/2

		midrow := mid / col
		midcol := mid % col
		midElement := matrix[midrow][midcol]

		if midElement == target {
			return true
		} else if midElement < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
