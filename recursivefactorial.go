package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	result := nb * RecursiveFactorial(nb-1)
	if nb > 1 && nb < 21 {
		return result
	}
	if result < 0 {
		return 0
	}
	return result
}
