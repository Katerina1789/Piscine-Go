package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= '0' && runes[i] <= '9') {
			if newWord {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] -= 32
				}
				newWord = false
			} else {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] += 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(runes)
}
