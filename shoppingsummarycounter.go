package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			if word != "" {
				result[word]++
				word = ""
			}
			// κάθε space προσθέτει 1 στο ""
			result[""]++
		} else {
			word += string(c)
		}
	}

	// Αν τελειώνει με λέξη
	if word != "" {
		result[word]++
	}

	return result
}
