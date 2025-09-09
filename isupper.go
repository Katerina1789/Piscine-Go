package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		} else if r >= 'a' && r <= 'z' {
			return false
		}
	}
	return false
}
