package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	upper := false

	if len(argument) > 0 && argument[0] == "--upper" {
		upper = true
		argument = argument[1:]
	}

	for _, i := range argument {
		n := 0

		for _, ch := range i {
			if ch < '0' || ch > '9' {
				n = 0
				break
			}
			n = n*10 + int(ch-'0')
		}

		if n >= 1 && n <= 26 {
			letter := 'a' + rune(n-1)
			if upper {
				letter = 'A' + rune(n-1)
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}
}
