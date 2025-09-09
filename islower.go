package piscine

func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'a' {
			return false
		}
	}
	return true
}
