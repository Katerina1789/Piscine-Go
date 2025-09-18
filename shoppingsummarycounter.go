package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		c := str[i]

		// if we see a space
		if c == ' ' {
			// only add if word is not empty
			if word != "" {
				result[word]++
				word = ""
			}
			// if word is empty, we just skip (this ignores multiple spaces)
		} else {
			word += string(c)
		}
	}

	// handle last word if it exists
	if word != "" {
		result[word]++
	}

	return result
}
