// You are given a 9 x 9 Sudoku board board. A Sudoku board is valid if the following rules are followed:

// Each row must contain the digits 1-9 without duplicates.
// Each column must contain the digits 1-9 without duplicates.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
// Return true if the Sudoku board is valid, otherwise return false

// Note: A board does not need to be full or be solvable to be valid.

package main

import "fmt"

func main() {

	board :=
		[][]byte{{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
			{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
			{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
			{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board))

}

func isValidSudoku(board [][]byte) bool {

	row := make([]map[byte]bool, 9)
	col := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]

			if val == '.' {
				continue
			}

			box := (r/3)*3 + (c / 3)

			if row[r][val] || col[c][val] || boxes[box][val] {
				return false
			}

			row[r][val] = true
			col[c][val] = true
			boxes[box][val] = true
		}
	}
	return true

}

// box := (r/3)*3 + (c/3)

// For each non-empty cell (r, c, val):
//   → 3 questions: seen in this row? this col? this box?
//   → If any yes → invalid
//   → If all no  → register in all three
