package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i > 0; i-- {
		printTwoDigit(i)
		z01.PrintRune(' ')
		printTwoDigit(i - 1)
		if i != 1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigit(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}
