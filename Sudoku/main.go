package main // Declares the main package — the entry point of the program

import (
	"os"
)

func main() { // Main function where execution begins
	if len(os.Args) != 10 {
		printError("Error: Invalid number of arguments") // If not, print an error message
		return                                           // Stops execution
	}

	var board [9][9]rune // Initializes a 9x9 Sudoku board using rune (character) type
	filledcells := 0     // Counter to track how many cells are filled

	for i := 0; i < 9; i++ { // Loops through each row
		row := os.Args[i+1] // Gets the i-th row from command-line arguments
		if len(row) != 9 {  // Checks if the row has exactly 9 characters
			printError("Error: Invalid number of characters") // Prints error if not
			return                                            // Stops execution
		}
		for j := 0; j < 9; j++ { // Loops through each character in the row
			if row[j] != '.' && (row[j] < '1' || row[j] > '9') { // Check if the character is valid
				printError("Error: Invalid character") // If invalid, print an error message
				return
			}
			board[i][j] = rune(row[j]) // Stores the character in the board
			if row[j] != '.' {
				filledcells++ // Increment the filled cell counter
			}
		}
	}

	if filledcells < 17 { // Checks if there are at least 17 filled cells
		printError("Error: Less than 17 filled cells") // Prints error if not
		return                                         // Stops execution
	}

	if hasDuplicates(board) { // Checks for duplicates in rows, columns, or boxes
		printError("Error: Duplicates found in row, column or 3x3 board") // Prints error if found
		return                                                            // Stops execution
	}

	if solve(&board) { // Attempts to solve the Sudoku puzzle
		printBoard(board) // If successful, prints the solved board
	} else {
		printError("Error: No solution") // If unsolvable, prints error
	}
}
