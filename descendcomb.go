package piscine

import "github.com/01-edu/z01"

func printNumber(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}

func DescendComb() {
	for i := 99; i >= 10; i-- {
		for j := i - 1; j >= 10; j-- {
			printNumber(i)
			z01.PrintRune(' ')
			printNumber(j)
			if !(i == 11 && j == 10) { // last pair is "11 10"
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
