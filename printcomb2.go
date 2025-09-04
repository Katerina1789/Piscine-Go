package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			z01.PrintRune(a)
			z01.PrintRune(b)
			if a == '9' && b == '9' {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
