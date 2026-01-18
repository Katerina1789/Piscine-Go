package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s { // loops through each character from string
		z01.PrintRune(r) // prints each character
	}
	z01.PrintRune('\n') // adds a newline at the end
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true // if there is a remainder it returns false meaning it is an odd number
	} else {
		return false // if the remainder is 0 it returns true meaning it is an even number
	}
}

func main() {
	argument := os.Args[1:]    // skips the name of the programme
	if isEven(len(argument)) { // if the length of the arguments is true in isEven function
		printStr("I have an even number of arguments") // prints this
	} else { // if the length of the arguments is false in the isEven function
		printStr("I have an odd number of arguments") // prints this
	}
}
