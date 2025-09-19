package piscine

func Rot14(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			rotated := r + 14
			if rotated > 'z' {
				rotated -= 26
			}
			result += string(rotated)
		} else if r >= 'A' && r <= 'Z' {
			rotated := r + 14
			if rotated > 'Z' {
				rotated -= 26
			}
			result += string(rotated)
		} else {
			result += string(r)
		}
	}
	return result
}
