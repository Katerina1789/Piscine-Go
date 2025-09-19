package piscine

import "github.com/01-edu/z01"

func printNumber(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printNumber(i)
			z01.PrintRune(' ')
			printNumber(j)
			if !(i == 1 && j == 0) { // last pair is "01 00"
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
