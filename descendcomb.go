package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	var result string
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			result += twoDigit(i) + " " + twoDigit(j)
			if !(i == 1 && j == 0) {
				result += ", "
			}
		}
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func twoDigit(n int) string {
	return string(rune(n/10+'0')) + string(rune(n%10+'0'))
}
