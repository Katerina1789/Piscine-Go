package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	word := ""
	spaceRun := 0
	sawWord := false

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			if word != "" {
				result[word]++
				word = ""
				sawWord = true
			}
			spaceRun++
		} else {
			if spaceRun > 0 {
				if sawWord {
					if spaceRun > 1 {
						result[""] += spaceRun - 1
					}
				} else {
					result[""] += spaceRun
				}
				spaceRun = 0
			}
			word += string(c)
		}
	}

	if word != "" {
		result[word]++
	}

	// εδώ το fix:
	if !sawWord && word == "" && len(str) > 0 {
		// string έχει μόνο spaces → μέτρησέ τα όλα
		result[""] = len(str)
	} else if spaceRun > 0 {
		// trailing spaces μετά από λέξη
		result[""] += spaceRun
	}

	return result
}
