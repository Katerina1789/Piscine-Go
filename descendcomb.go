package piscine

import (
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printStr(twoDigit(i) + " " + twoDigit(j))
			if !(i == 1 && j == 0) {
				printStr(", ")
			}
		}
	}
	z01.PrintRune('\n')
}

func twoDigit(n int) string {
	return string(rune(n/10+'0')) + string(rune(n%10+'0'))
}
