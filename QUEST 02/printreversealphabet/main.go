package main

import "github.com/01-edu/z01"

func main() {
	ab := "zyxwvutsrqponmlkjihgfedcba"
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(ab[i]))
	}
	z01.PrintRune('\n')
}
