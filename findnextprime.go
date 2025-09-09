package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		nb = 2
	}
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
