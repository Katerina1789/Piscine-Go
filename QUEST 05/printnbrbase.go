package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	baseLen := len(base)
	var result []rune
	for nbr > 0 {
		remainder := nbr % baseLen
		result = append(result, rune(base[remainder]))
		nbr = nbr / baseLen
	}
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}
