package piscine

func ToUpper(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' || r <= 'z' {
			return result + string(r-32)
		}
	}
	return result
}
