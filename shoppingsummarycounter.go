package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""
	spaceCount := 0

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			// Αν υπάρχει ήδη λέξη, την αποθηκεύουμε και ξεκινάμε νέα
			if word != "" {
				result[word]++
				word = ""
			} else {
				// Διαδοχικά ή αρχικά κενά -> μετράμε ως extra spaces
				spaceCount++
			}
		} else {
			// Αν υπήρχαν spaces πριν από γράμμα, τα καταχωρούμε πρώτα
			if spaceCount > 0 {
				result[""] += spaceCount
				spaceCount = 0
			}
			word += string(c)
		}
	}

	// Αν τελειώνει με λέξη
	if word != "" {
		result[word]++
	}

	// Αν τελειώνει με spaces
	if spaceCount > 0 {
		result[""] += spaceCount
	}

	return result
}
