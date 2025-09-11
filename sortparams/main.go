package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for i := 1; i < len(argument)-1; i++ {
		for j := i + 1; j < len(argument); j++ {
			if argument[i] > argument[j] {
				argument[i], argument[j] = argument[j], argument[i]
			}
		}
	}
	for i := 1; i < len(argument); i++ {
		for _, ch := range argument[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
