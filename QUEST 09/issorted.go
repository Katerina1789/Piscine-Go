package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascending = false
		}
		if f(a[i], a[i+1]) < 0 {
			descending = false
		}
	}

	return ascending || descending
}
