package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		} else if r < 'A' || r > 'Z' {
			return false
		} else if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
