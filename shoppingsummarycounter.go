package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	word := ""
	spaceRun := 0    // length of current consecutive spaces
	sawWord := false // have we completed at least one word yet?

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			// finalize a word if we were building one
			if word != "" {
				result[word]++
				word = ""
				sawWord = true
			}
			spaceRun++
		} else {
			// we’re entering a word after some spaces: account for spaces first
			if spaceRun > 0 {
				if sawWord {
					// between words: only count the extras beyond the first separator
					if spaceRun > 1 {
						result[""] += spaceRun - 1
					}
				} else {
					// leading spaces: count them all
					result[""] += spaceRun
				}
				spaceRun = 0
			}
			// build the current word
			word += string(c)
		}
	}

	// finalize last word
	if word != "" {
		result[word]++
	}

	// trailing spaces: count them all
	if spaceRun > 0 {
		result[""] += spaceRun
	}

	return result
}
