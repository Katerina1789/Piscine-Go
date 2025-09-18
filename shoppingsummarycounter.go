package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	// Step 1: Split the string manually into words
	word := ""
	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			if word != "" {
				result[word]++
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	// Step 2: Add the last word (if not followed by space)
	if word != "" {
		result[word]++
	}

	return result
}
