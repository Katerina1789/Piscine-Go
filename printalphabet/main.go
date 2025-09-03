package main

import "github.com/01-edu/z01"

func main() {
	ab := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 25; i++ {
		z01.PrintRune(rune(ab[i]))
	}
	z01.PrintRune('\n')
}
