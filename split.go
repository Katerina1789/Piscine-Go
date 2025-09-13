package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	result := make([]string, 0)
	word := ""
	i := 0

	for i < len(s) {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, word)
			word = ""
			i += len(sep)
		} else {
			word += string(s[i])
			i++
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
