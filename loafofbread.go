package piscine

func LoafOfBread(str string) string {
	// Remove spaces from the input string
	cleaned := ""
	for _, r := range str {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	if len(cleaned) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(cleaned); {
		// Collect 5 characters
		word := ""
		for j := 0; j < 5 && i < len(cleaned); j++ {
			word += string(cleaned[i])
			i++
		}
		result += word

		// Skip the next character if it exists
		if i < len(cleaned) {
			i++
		}

		result += " "
		count++
	}

	// Remove trailing space and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return result + "\n"
}
