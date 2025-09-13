package main // Declares the main package — the entry point of the program

import "github.com/01-edu/z01" // Imports the z01 library for printing runes (characters) to the terminal

// Prints an error message character by character
func printError(msg string) {
	for _, char := range msg { // Loop through each character in the string
		z01.PrintRune(char) // Print each character individually
	}
	z01.PrintRune('\n') // Add a newline after the message
}

// Prints the Sudoku board in a formatted way
func printBoard(board [9][9]rune) {
	for row := 0; row < 9; row++ { // Loop through each row
		for col := 0; col < 9; col++ { // Loop through each column
			z01.PrintRune(board[row][col]) // Print the cell value
			z01.PrintRune(' ')             // Add a space between numbers
		}
		z01.PrintRune('\n') // Print a newline after each row
	}
}

// Checks for duplicates in rows, columns, and 3x3 boxes
func hasDuplicates(board [9][9]rune) bool {
	for i := 0; i < 9; i++ { // Loop through each row and column index
		row := [9]bool{} // Tracks seen digits in the current row
		col := [9]bool{} // Tracks seen digits in the current column
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' { // Skip empty cells
				n := board[i][j] - '1' // Convert rune to index (0–8)
				if row[n] {            // If already seen in row
					return true // Duplicate found
				}
				row[n] = true // Mark as seen
			}
			if board[j][i] != '.' { // Same for column
				n := board[j][i] - '1'
				if col[n] {
					return true
				}
				col[n] = true
			}
		}
	}

	// Check each 3x3 box
	for row := 0; row < 9; row += 3 { // Loop through box starting rows
		for col := 0; col < 9; col += 3 { // Loop through box starting columns
			box := [9]bool{} // Tracks seen digits in the current box
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					if board[i][j] != '.' {
						n := board[i][j] - '1'
						if box[n] {
							return true
						}
						box[n] = true
					}
				}
			}
		}
	}
	return false // No duplicates found
}

// Checks if a number can be placed in a specific cell
func isValid(board [9][9]rune, row, col int, num rune) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false // Number already exists in row or column
		}
	}
	startRow := (row / 3) * 3 // Starting row of the 3x3 box
	startCol := (col / 3) * 3 // Starting column of the 3x3 box
	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			if board[row][col] == num {
				return false // Number already exists in box
			}
		}
	}
	return true // Number is valid in this cell
}

// Solves the Sudoku puzzle using backtracking
func solve(board *[9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' { // If the current cell is empty
				for num := '1'; num <= '9'; num++ { // Try all possible numbers from 1 to 9
					if isValid(*board, row, col, num) { // Check if the number is valid in this cell
						board[row][col] = num // Place the number in the cell
						if solve(board) {     // Recursively try to solve the rest of the puzzle
							return true // If the board is solved, return true
						}
						board[row][col] = '.' // If the number doesn't work, reset the cell and try next
					}
				}
				return false // If no number can be placed in this cell, return false
			}
		}
	}
	return true // If all cells are filled, return true (the board is solved)
}
