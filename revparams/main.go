package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for i := len(argument); i < len(argument); i-- {
		for _, ch := range argument[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
